package services

import (
	consulapi "github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
)

type ConsulService struct{}

func (cs *ConsulService) Start() error {
	return nil
}

func (cs *ConsulService) Stop() error {
	return nil
}

func newConsulService() *ConsulService {
	return &ConsulService{}
}

func newClient() error {
	c, _ := consulapi.NewClient(consulapi.DefaultConfig())
	logrus.Info(c)
	return nil
}

func (cs *ConsulService) AddNode() error {
	return nil
}

func (cs *ConsulService) RemoveNode() error {
	return nil
}
