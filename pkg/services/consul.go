package services

type ConsulService struct{}

func (cs *ConsulService) Start() error {
	return nil
}

func (cs *ConsulService) Stop() error {
	return nil
}

func NewConsulService() *ConsulService {
	return &ConsulService{}
}
