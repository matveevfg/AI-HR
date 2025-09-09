package aiHr

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

type FileTransfer struct {
	file    *os.File
	mutex   sync.Mutex
	chunks  map[int][]byte
	nextID  int
	tempDir string
}

func (s *Service) ReadAudio(ctx context.Context, ws *websocket.Conn) {
	ft := &FileTransfer{
		tempDir: os.TempDir(),
	}

	defer func() {
		ft.cleanup()
	}()

	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			break
		}

		if messageType == websocket.BinaryMessage {
			if err := ft.startNewFile(); err != nil {
				_ = ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("ERROR: %v", err)))

				continue
			}

			err := ft.handleFileChunk(message, ws)
			if err != nil {
				_ = ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("ERROR: %v", err)))

				continue
			}

			filePath, err := ft.finalizeFile()
			if err != nil {
				_ = ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("ERROR: %v", err)))
			}

			println(filePath)
		}
	}
}

func (ft *FileTransfer) startNewFile() error {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()

	tempFile, err := os.CreateTemp(ft.tempDir, "upload_*.tmp")
	if err != nil {
		return fmt.Errorf("не удалось создать временный файл: %v", err)
	}

	ft.file = tempFile
	ft.chunks = make(map[int][]byte)
	ft.nextID = 0

	return nil
}

func (ft *FileTransfer) handleFileChunk(chunk []byte, ws *websocket.Conn) error {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()

	if ft.file == nil {
		return fmt.Errorf("файл не инициализирован, отправьте команду START")
	}

	_, err := ft.file.Write(chunk)
	if err != nil {
		return fmt.Errorf("ошибка записи чанка в файл: %v", err)
	}

	_ = ws.WriteMessage(websocket.TextMessage, []byte("ok"))

	return nil
}

func (ft *FileTransfer) finalizeFile() (string, error) {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()

	if ft.file == nil {
		return "", fmt.Errorf("файл не инициализирован")
	}

	err := ft.file.Close()
	if err != nil {
		return "", fmt.Errorf("ошибка закрытия файла: %v", err)
	}

	return ft.file.Name(), nil
}

func (ft *FileTransfer) cleanup() {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()

	if ft.file != nil {
		_ = ft.file.Close()
		_ = os.Remove(ft.file.Name())
	}
}
