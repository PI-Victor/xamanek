package core

import (
	"github.com/sirupsen/logrus"

	"github.com/cloudflavor/xamanek/pkg/services"
)

type Config struct {
	LogLevel         logrus.Level `json:"logLevel"`
	ContainerRuntime string       `json:"containerRuntime"`
	IsMaster         bool         `json:"isMaster"`
}

func NewConfig() *Config {
	return &Config{}
}

type Instance struct {
	Services []services.Service
}

func (i *Instance) Start() error {
	for _, service := range i.Services {
		service.Start()
		logrus.Info(service)
	}
	return nil
}

func NewInstance(config *Config) *Instance {
	consulService := services.NewConsulService()

	s := append([]services.Service{}, consulService)
	return &Instance{
		Services: s,
	}
}
