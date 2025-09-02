package aiHr

type Service struct {
	storage storage
}

func New(storage storage) *Service {
	return &Service{
		storage: storage,
	}
}
