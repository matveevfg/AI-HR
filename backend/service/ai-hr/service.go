package aiHr

type Service struct {
	storage             storage
	llmClient           llmClient
	transcriptionClient transcriptionClient
}

func New(storage storage, llmClient llmClient, transcriptionClient transcriptionClient) *Service {
	return &Service{
		storage:             storage,
		llmClient:           llmClient,
		transcriptionClient: transcriptionClient,
	}
}
