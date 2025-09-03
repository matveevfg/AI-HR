package deepseek

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-deepseek/deepseek"
	"github.com/go-deepseek/deepseek/request"
	"github.com/go-deepseek/deepseek/response"
)

type Service struct {
	storage        Storage
	client         *http.Client
	deepseekClient deepseek.Client
}

func New(storage Storage, token string) (*Service, error) {
	client, err := deepseek.NewClient(token)
	if err != nil {
		return nil, err
	}

	return &Service{
		storage: storage,
		client: &http.Client{
			Timeout: defaultTimeout,
		},
		deepseekClient: client,
	}, nil
}

func (s *Service) createChatCompletion(ctx context.Context, maxTokens int, messages []*request.Message) (*response.ChatCompletionsResponse, error) {
	resp, err := s.deepseekClient.CallChatCompletionsChat(
		ctx,
		&request.ChatCompletionsRequest{
			MaxTokens: maxTokens,
			Model:     deepseek.DEEPSEEK_CHAT_MODEL,
			Messages:  messages,
		},
	)
	if err != nil {
		return nil, err
	}

	if len(resp.Choices) == 0 {
		return nil, errors.New("response is empty")
	}

	return resp, nil
}
