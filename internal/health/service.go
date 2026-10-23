package health

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetStatus() string {
	return "ok"
}
