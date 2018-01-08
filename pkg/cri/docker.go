package cri

type DockerService struct{}

func (ds *DockerService) Start() error {
	return nil
}

func (ds *DockerService) Stop() error {
	return nil
}

func NewDockerService() *DockerService {
	return &DockerService{}
}
