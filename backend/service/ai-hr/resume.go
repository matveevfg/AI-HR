package aiHr

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/lu4p/cat"

	"github.com/matveevfg/AI-HR/backend/models"
)

func (s *Service) SaveResume(ctx context.Context, files []*multipart.FileHeader, vacancyID uuid.UUID) error {
	ctx, err := s.storage.CtxWithTx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = s.storage.TxRollback(ctx)
	}()

	for _, file := range files {
		if !strings.HasSuffix(strings.ToLower(file.Filename), ".rtf") {
			return errors.New("resume field must end with .rtf")
		}

		src, err := file.Open()
		if err != nil {
			return err
		}

		fileData, err := io.ReadAll(src)
		if err != nil {
			return err
		}

		tmp, err := os.Create(os.TempDir() + "/" + file.Filename)
		if err != nil {
			return err
		}

		if err := os.WriteFile(tmp.Name(), fileData, 0666); err != nil {
			return err
		}

		if err := convertRTFToDOCX(tmp.Name(), "/tmp/out.docx"); err != nil {
			return err
		}

		text, err := extractTextFromDocx("/tmp/out.docx")
		if err != nil {
			return err
		}

		resume, err := s.llmClient.ResumeToJSON(ctx, text)
		if err != nil {
			return err
		}

		resume.ID = uuid.New()
		resume.VacancyID = vacancyID

		if err := s.storage.SaveWorkPlaces(ctx, resume.WorkPlaces); err != nil {
			return err
		}

		if err := s.storage.SaveResume(ctx, resume); err != nil {
			return err
		}

		if err := src.Close(); err != nil {
			return err
		}

		if err := os.Remove(tmp.Name()); err != nil {
			return err
		}

		if err := os.Remove("/tmp/out.docx"); err != nil {
			return err
		}
	}

	if err := s.storage.TxCommit(ctx); err != nil {
		return err
	}

	return nil
}

func extractTextFromDocx(filePath string) (string, error) {
	txt, err := cat.File(filePath)
	if err != nil {
		return "", err
	}

	return txt, nil
}

func convertRTFToDOCX(inputPath, outputPath string) error {
	cmd := exec.Command("libreoffice",
		"--headless",
		"--convert-to",
		"docx",
		inputPath,
		"--outdir",
		filepath.Dir(outputPath),
	)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("conversion failed: %v", err)
	}

	defaultOutput := strings.TrimSuffix(inputPath, filepath.Ext(inputPath)) + ".docx"

	if defaultOutput != outputPath {
		if err := os.Rename(defaultOutput, outputPath); err != nil {
			return fmt.Errorf("failed to rename file: %v", err)
		}
	}

	return nil
}

func (s *Service) Resumes(ctx context.Context, vacancyID uuid.UUID) ([]*models.Resume, error) {
	return s.storage.Resumes(ctx, vacancyID)
}
