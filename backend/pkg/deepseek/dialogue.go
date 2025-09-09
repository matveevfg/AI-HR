package deepseek

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-deepseek/deepseek/request"
	"github.com/sashabaranov/go-openai"
)

func (s *Service) VacancyInterviewPlan(ctx context.Context, vacancy string) (string, error) {
	var msgs []*request.Message
	msgs = append(msgs, &request.Message{
		Role:    openai.ChatMessageRoleUser,
		Content: fmt.Sprintf(interviewPlanPrompt, vacancy),
	})

	resp, err := s.createChatCompletion(ctx, maxTokens, msgs)
	if err != nil {
		return "", fmt.Errorf("create chat completion: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("response is empty")
	}

	return resp.Choices[0].Message.Content, nil
}

func (s *Service) GenerateDialogue(ctx context.Context, plan, dialogue string) (string, error) {
	var msgs []*request.Message
	msgs = append(msgs, &request.Message{
		Role:    openai.ChatMessageRoleUser,
		Content: fmt.Sprintf(dialogPrompt, plan, dialogue),
	})

	resp, err := s.createChatCompletion(ctx, maxTokens, msgs)
	if err != nil {
		return "", fmt.Errorf("create chat completion: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("response is empty")
	}

	return resp.Choices[0].Message.Content, nil
}
