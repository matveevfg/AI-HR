package transcription

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Audio constants matching the T-one pipeline
	SampleRate     = 8000 // 8kHz
	ChunkSize      = 2400 // 300ms at 8kHz
	BytesPerSample = 2    // 16-bit audio
	ChunkSizeBytes = ChunkSize * BytesPerSample
)

// APIResponse represents the JSON responses from the server
type APIResponse struct {
	Event     string  `json:"event"`
	Message   string  `json:"message,omitempty"`
	Text      string  `json:"text,omitempty"`
	StartTime float64 `json:"start_time,omitempty"`
	EndTime   float64 `json:"end_time,omitempty"`
	IsFinal   bool    `json:"is_final,omitempty"`
}

// Client represents the T-one ASR WebSocket client
type Client struct {
	conn   *websocket.Conn
	url    string
	logger *log.Logger
}

// NewClient creates a new T-one ASR client
func NewClient(serverURL string) *Client {
	return &Client{
		url:    serverURL,
		logger: log.New(os.Stdout, "[T-one Client] ", log.LstdFlags),
	}
}

// Connect establishes WebSocket connection to the server
func (c *Client) Connect(ctx context.Context) error {
	u, err := url.Parse(c.url)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}

	c.logger.Printf("Connecting to %s", c.url)

	dialer := websocket.DefaultDialer
	conn, _, err := dialer.DialContext(ctx, u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}

	c.conn = conn
	c.logger.Println("Connected successfully")
	return nil
}

// Close closes the WebSocket connection
func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// SendAudioFile sends a WAV file for transcription
func (c *Client) SendAudioFile(ctx context.Context, filePath string) error {
	if c.conn == nil {
		return fmt.Errorf("not connected")
	}

	audioData, err := c.readWAVFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read WAV file: %w", err)
	}

	c.logger.Printf("Sending audio file: %s (%d samples)", filePath, len(audioData))

	responseChan := make(chan APIResponse, 100)
	errorChan := make(chan error, 1)

	go c.listenForResponses(responseChan, errorChan)

	select {
	case response := <-responseChan:
		if response.Event != "ready" {
			return fmt.Errorf("expected ready signal, got: %s", response.Event)
		}
		c.logger.Println("Server is ready, starting audio transmission")
	case err := <-errorChan:
		return fmt.Errorf("error waiting for ready signal: %w", err)
	case <-ctx.Done():
		return ctx.Err()
	}

	totalChunks := (len(audioData) + ChunkSize - 1) / ChunkSize

	for i := 0; i < len(audioData); i += ChunkSize {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		end := i + ChunkSize
		if end > len(audioData) {
			end = len(audioData)
		}

		chunk := audioData[i:end]

		if len(chunk) < ChunkSize {
			paddedChunk := make([]int16, ChunkSize)
			copy(paddedChunk, chunk)
			chunk = paddedChunk
		}

		chunkBytes := make([]byte, len(chunk)*2)
		for j, sample := range chunk {
			binary.LittleEndian.PutUint16(chunkBytes[j*2:], uint16(sample))
		}

		if err := c.conn.WriteMessage(websocket.BinaryMessage, chunkBytes); err != nil {
			return fmt.Errorf("failed to send audio chunk: %w", err)
		}

		chunkNum := (i / ChunkSize) + 1
		c.logger.Printf("Sent chunk %d/%d", chunkNum, totalChunks)

		if i+ChunkSize < len(audioData) {
			select {
			case response := <-responseChan:
				if response.Event == "transcription" {
					c.logger.Printf("Transcription: %s", response.Text)
				}
			case err := <-errorChan:
				return fmt.Errorf("error during transmission: %w", err)
			case <-time.After(5 * time.Second):
				c.logger.Println("Warning: No response received, continuing...")
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}

	if err := c.conn.WriteMessage(websocket.BinaryMessage, []byte{}); err != nil {
		return fmt.Errorf("failed to send end signal: %w", err)
	}

	c.logger.Println("Audio transmission completed, waiting for final results...")

	for {
		select {
		case response := <-responseChan:
			switch response.Event {
			case "transcription":
				c.logger.Printf("Final transcription: %s", response.Text)
			case "complete":
				c.logger.Println("Transcription completed successfully")
				return nil
			case "error":
				return fmt.Errorf("server error: %s", response.Message)
			}
		case err := <-errorChan:
			return fmt.Errorf("error waiting for completion: %w", err)
		case <-time.After(10 * time.Second):
			return fmt.Errorf("timeout waiting for completion")
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// SendSyntheticAudio sends generated sine wave audio for testing
func (c *Client) SendSyntheticAudio(ctx context.Context, durationSeconds float64, frequency float64) error {
	if c.conn == nil {
		return fmt.Errorf("not connected")
	}

	numSamples := int(durationSeconds * SampleRate)
	audioData := make([]int16, numSamples)

	for i := 0; i < numSamples; i++ {
		t := float64(i) / SampleRate
		sample := math.Sin(2 * math.Pi * frequency * t)
		audioData[i] = int16(sample * 16383)
	}

	c.logger.Printf("Sending synthetic audio: %.1fs at %.0fHz (%d samples)",
		durationSeconds, frequency, numSamples)

	responseChan := make(chan APIResponse, 100)
	errorChan := make(chan error, 1)

	go c.listenForResponses(responseChan, errorChan)

	select {
	case response := <-responseChan:
		if response.Event != "ready" {
			return fmt.Errorf("expected ready signal, got: %s", response.Event)
		}
		c.logger.Println("Server is ready, starting synthetic audio transmission")
	case err := <-errorChan:
		return fmt.Errorf("error waiting for ready signal: %w", err)
	case <-ctx.Done():
		return ctx.Err()
	}

	return c.sendAudioData(ctx, audioData, responseChan, errorChan)
}

// sendAudioData sends audio data in chunks
func (c *Client) sendAudioData(ctx context.Context, audioData []int16, responseChan chan APIResponse, errorChan chan error) error {
	totalChunks := (len(audioData) + ChunkSize - 1) / ChunkSize

	for i := 0; i < len(audioData); i += ChunkSize {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		end := i + ChunkSize
		if end > len(audioData) {
			end = len(audioData)
		}

		chunk := audioData[i:end]

		if len(chunk) < ChunkSize {
			paddedChunk := make([]int16, ChunkSize)
			copy(paddedChunk, chunk)
			chunk = paddedChunk
		}

		chunkBytes := make([]byte, len(chunk)*2)
		for j, sample := range chunk {
			binary.LittleEndian.PutUint16(chunkBytes[j*2:], uint16(sample))
		}

		if err := c.conn.WriteMessage(websocket.BinaryMessage, chunkBytes); err != nil {
			return fmt.Errorf("failed to send audio chunk: %w", err)
		}

		chunkNum := (i / ChunkSize) + 1
		c.logger.Printf("Sent chunk %d/%d", chunkNum, totalChunks)

		time.Sleep(300 * time.Millisecond)
	}

	if err := c.conn.WriteMessage(websocket.BinaryMessage, []byte{}); err != nil {
		return fmt.Errorf("failed to send end signal: %w", err)
	}

	c.logger.Println("Audio transmission completed")
	return nil
}

// listenForResponses listens for server responses
func (c *Client) listenForResponses(responseChan chan<- APIResponse, errorChan chan<- error) {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			errorChan <- err
			return
		}

		var response APIResponse
		if err := json.Unmarshal(message, &response); err != nil {
			errorChan <- fmt.Errorf("failed to parse response: %w", err)
			return
		}

		responseChan <- response
	}
}

// readWAVFile reads a WAV file and returns 16-bit PCM audio data
func (c *Client) readWAVFile(filePath string) ([]int16, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer func() {
		_ = file.Close()
	}()

	header := make([]byte, 44)
	if _, err := io.ReadFull(file, header); err != nil {
		return nil, fmt.Errorf("failed to read WAV header: %w", err)
	}

	if string(header[0:4]) != "RIFF" || string(header[8:12]) != "WAVE" {
		return nil, fmt.Errorf("not a valid WAV file")
	}

	audioFormat := binary.LittleEndian.Uint16(header[20:22])
	numChannels := binary.LittleEndian.Uint16(header[22:24])
	sampleRate := binary.LittleEndian.Uint32(header[24:28])
	bitsPerSample := binary.LittleEndian.Uint16(header[34:36])

	c.logger.Printf("WAV file info: %dHz, %d channels, %d bits", sampleRate, numChannels, bitsPerSample)

	if audioFormat != 1 {
		return nil, fmt.Errorf("only PCM format supported (got format %d)", audioFormat)
	}
	if bitsPerSample != 16 {
		return nil, fmt.Errorf("only 16-bit audio supported (got %d bits)", bitsPerSample)
	}

	audioBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read audio data: %w", err)
	}

	audioData := make([]int16, len(audioBytes)/2)
	for i := 0; i < len(audioData); i++ {
		audioData[i] = int16(binary.LittleEndian.Uint16(audioBytes[i*2:]))
	}

	if numChannels == 2 {
		c.logger.Println("Converting stereo to mono")
		monoData := make([]int16, len(audioData)/2)
		for i := 0; i < len(monoData); i++ {
			left := audioData[i*2]
			right := audioData[i*2+1]
			monoData[i] = int16((int32(left) + int32(right)) / 2)
		}
		audioData = monoData
	}

	if sampleRate != SampleRate {
		c.logger.Printf("Resampling from %dHz to %dHz", sampleRate, SampleRate)
		audioData = c.resample(audioData, int(sampleRate), SampleRate)
	}

	return audioData, nil
}

// resample performs basic audio resampling
func (c *Client) resample(input []int16, fromRate, toRate int) []int16 {
	if fromRate == toRate {
		return input
	}

	ratio := float64(fromRate) / float64(toRate)
	outputLen := int(float64(len(input)) / ratio)
	output := make([]int16, outputLen)

	for i := 0; i < outputLen; i++ {
		srcIndex := float64(i) * ratio
		srcIndexInt := int(srcIndex)

		if srcIndexInt >= len(input)-1 {
			output[i] = input[len(input)-1]
		} else {
			frac := srcIndex - float64(srcIndexInt)
			sample1 := float64(input[srcIndexInt])
			sample2 := float64(input[srcIndexInt+1])
			interpolated := sample1 + frac*(sample2-sample1)
			output[i] = int16(interpolated)
		}
	}

	return output
}
