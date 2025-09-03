package aiHr

type Service struct {
	storage   storage
	llmClient llmClient
}

func New(storage storage, llmClient llmClient) *Service {
	return &Service{
		storage:   storage,
		llmClient: llmClient,
	}
}
