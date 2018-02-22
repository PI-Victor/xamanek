package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func newServices() []*api.AgentServiceRegistration {
	/**
	  type AgentServiceRegistration struct {
	  ID                string   `json:",omitempty"`
	  Name              string   `json:",omitempty"`
	  Tags              []string `json:",omitempty"`
	  Port              int      `json:",omitempty"`
	  Address           string   `json:",omitempty"`
	  EnableTagOverride bool     `json:",omitempty"`
	}
	**/
	return []*api.AgentServiceRegistration{
		{
			ID:   "newAgent",
			Name: "test-agent",
		},
		{
			ID:   "newAgent2",
			Name: "test-agen2",
		},
		{
			ID:   "newAgent3",
			Name: "test-agent3",
		},
		{
			ID:   "newAgent4",
			Name: "test-agent4",
		},
		{
			ID:   "newAgent5",
			Name: "test-agen5",
		},
		{
			ID:   "newAgent6",
			Name: "test-agent6",
		},
	}
}

func main() {
	newSvcs := newServices()
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	c := client.Agent()
	for _, svc := range newSvcs {
		c.ServiceRegister(svc)
	}

	svc, err := c.Services()
	if err != nil {
		panic(err)
	}

	for _, s := range svc {
		fmt.Printf("%#v", s)
	}

}
