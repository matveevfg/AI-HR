package textToAudio

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Service struct{}

const (
	maxChunkLength = 200 // Максимальная длина текста для одного запроса
)

func (s *Service) TextToSpeech(text, language, outputFile string) error {
	// Разбиваем длинный текст на части
	chunks := splitText(text, maxChunkLength)

	// Создаем временные файлы для каждого чанка
	var tempFiles []string

	for i, chunk := range chunks {
		if strings.TrimSpace(chunk) == "" {
			continue
		}

		tempFile := fmt.Sprintf("temp_%d.mp3", i)
		err := googleTTS(chunk, language, tempFile)
		if err != nil {
			// Удаляем временные файлы при ошибке
			cleanupTempFiles(tempFiles)
			return fmt.Errorf("ошибка в чанке %d: %v", i, err)
		}
		tempFiles = append(tempFiles, tempFile)
	}

	// Объединяем аудиофайлы
	err := mergeAudioFiles(tempFiles, outputFile)
	cleanupTempFiles(tempFiles)

	return err
}

func splitText(text string, maxLength int) []string {
	var chunks []string
	words := strings.Fields(text)

	currentChunk := ""
	for _, word := range words {
		if len(currentChunk)+len(word)+1 > maxLength {
			chunks = append(chunks, currentChunk)
			currentChunk = word
		} else {
			if currentChunk != "" {
				currentChunk += " " + word
			} else {
				currentChunk = word
			}
		}
	}

	if currentChunk != "" {
		chunks = append(chunks, currentChunk)
	}

	return chunks
}

func googleTTS(text, language, outputFile string) error {
	encodedText := url.QueryEscape(text)
	ttsURL := fmt.Sprintf("https://translate.google.com/translate_tts?ie=UTF-8&tl=%s&client=tw-ob&q=%s",
		language, encodedText)

	client := &http.Client{}
	req, err := http.NewRequest("GET", ttsURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Referer", "https://translate.google.com/")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP ошибка: %s", resp.Status)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func mergeAudioFiles(inputFiles []string, outputFile string) error {
	// В реальном проекте лучше использовать ffmpeg или другую библиотеку
	// Здесь простейшее объединение (для MP3 это будет работать некорректно)

	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	for _, inputFile := range inputFiles {
		input, err := os.Open(inputFile)
		if err != nil {
			return err
		}

		_, err = io.Copy(output, input)
		input.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func cleanupTempFiles(files []string) {
	for _, file := range files {
		os.Remove(file)
	}
}
