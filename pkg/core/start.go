package core

import (
	"github.com/sirupsen/logrus"

	"github.com/cloudflavor/xamanek/pkg/services"
)

type Config struct {
	LogLevel         logrus.Level `json:"logLevel"`
	ContainerRuntime string       `json:"containerRuntime"`
	IsMaster         bool         `json:"isMaster"`
	MasterURI        string       `json:"masterURI,omitempty"`
}

func NewConfig() *Config {
	return &Config{}
}

type Server struct {
	Services []services.Service
}

func (i *Server) Start() error {
	for _, service := range i.Services {
		service.Start()
		logrus.Info(service)
	}
	return nil
}

func NewInstance(config *Config) *Server {
	return nil
}
