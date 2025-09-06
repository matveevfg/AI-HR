package deepseek

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/go-deepseek/deepseek/request"
	"github.com/matveevfg/AI-HR/backend/models"
	"github.com/sashabaranov/go-openai"
)

func (s *Service) VacancyToJSON(ctx context.Context, vacancyText string) (*models.Vacancy, error) {
	var msgs []*request.Message
	msgs = append(msgs, &request.Message{
		Role:    openai.ChatMessageRoleUser,
		Content: fmt.Sprintf(vacancyToJSONPrompt, vacancyText),
	})

	resp, err := s.createChatCompletion(ctx, maxTokens, msgs)
	if err != nil {
		return nil, fmt.Errorf("create chat completion: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, errors.New("response is empty")
	}

	var vacancy models.Vacancy
	if err := json.Unmarshal([]byte(clearJSON(resp.Choices[0].Message.Content)), &vacancy); err != nil {
		return nil, err
	}

	return &vacancy, nil
}

func clearJSON(msg string) string {
	msg = strings.Replace(msg, "```json", "", -1)
	msg = strings.Replace(msg, "```", "", -1)

	return msg
}
