package deepseek

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-deepseek/deepseek/request"
	"github.com/sashabaranov/go-openai"

	"github.com/matveevfg/AI-HR/backend/models"
)

func (s *Service) ResumeToJSON(ctx context.Context, resumeText string) (*models.Resume, error) {
	var msgs []*request.Message
	msgs = append(msgs, &request.Message{
		Role:    openai.ChatMessageRoleUser,
		Content: fmt.Sprintf(resumeToJSONPrompt, resumeText),
	})

	resp, err := s.createChatCompletion(ctx, maxTokens, msgs)
	if err != nil {
		return nil, fmt.Errorf("create chat completion: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, errors.New("response is empty")
	}

	var resume models.Resume
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &resume); err != nil {
		return nil, err
	}

	return &resume, nil
}
