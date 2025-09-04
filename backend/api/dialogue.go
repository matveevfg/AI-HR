package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type FileTransfer struct {
	file    *os.File
	mutex   sync.Mutex
	chunks  map[int][]byte
	nextID  int
	tempDir string
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Разрешаем подключения с любых источников (для разработки)
	},
}

func (s *Server) handleWebSocket(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	ft := &FileTransfer{
		chunks:  make(map[int][]byte),
		nextID:  0,
		tempDir: os.TempDir(),
	}

	defer func() {
		_ = ws.Close()
		ft.cleanup()
	}()

	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			break
		}

		if messageType == websocket.BinaryMessage {
			err := ft.handleFileChunk(message, ws)
			if err != nil {
				_ = ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("ERROR: %v", err)))

				continue
			}
		} else if messageType == websocket.TextMessage {
			command := string(message)
			switch command {
			case "START":
				err := ft.startNewFile()
				if err != nil {
					_ = ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("ERROR: %v", err)))
				} else {
					_ = ws.WriteMessage(websocket.TextMessage, []byte("READY"))
				}
			case "END":
				err := ft.finalizeFile()
				if err != nil {
					_ = ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("ERROR: %v", err)))
				} else {
					_ = ws.WriteMessage(websocket.TextMessage, []byte("COMPLETE"))
				}
			default:
				_ = ws.WriteMessage(websocket.TextMessage, []byte("UNKNOWN_COMMAND"))
			}
		}
	}

	return nil
}

func (ft *FileTransfer) startNewFile() error {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()

	// Создаем временный файл
	tempFile, err := os.CreateTemp(ft.tempDir, "upload_*.tmp")
	if err != nil {
		return fmt.Errorf("не удалось создать временный файл: %v", err)
	}

	ft.file = tempFile
	ft.chunks = make(map[int][]byte)
	ft.nextID = 0

	log.Printf("Создан временный файл: %s", tempFile.Name())
	return nil
}

func (ft *FileTransfer) handleFileChunk(chunk []byte, ws *websocket.Conn) error {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()

	if ft.file == nil {
		return fmt.Errorf("файл не инициализирован, отправьте команду START")
	}

	// Сохраняем чанк
	ft.chunks[ft.nextID] = chunk
	ft.nextID++

	// Записываем чанк в файл
	_, err := ft.file.Write(chunk)
	if err != nil {
		return fmt.Errorf("ошибка записи чанка в файл: %v", err)
	}

	// Отправляем подтверждение клиенту
	ack := fmt.Sprintf("ACK:%d", ft.nextID-1)
	_ = ws.WriteMessage(websocket.TextMessage, []byte(ack))

	log.Printf("Получен чанк #%d, размер: %d байт", ft.nextID-1, len(chunk))
	return nil
}

func (ft *FileTransfer) finalizeFile() error {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()

	if ft.file == nil {
		return fmt.Errorf("файл не инициализирован")
	}

	// Закрываем файл
	err := ft.file.Close()
	if err != nil {
		return fmt.Errorf("ошибка закрытия файла: %v", err)
	}

	fileInfo, err := os.Stat(ft.file.Name())
	if err != nil {
		return fmt.Errorf("ошибка получения информации о файле: %v", err)
	}

	log.Printf("Файл успешно сохранен: %s, размер: %d байт",
		filepath.Base(ft.file.Name()), fileInfo.Size())

	return nil
}

func (ft *FileTransfer) cleanup() {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()

	if ft.file != nil {
		_ = ft.file.Close()
		_ = os.Remove(ft.file.Name())
	}
}
