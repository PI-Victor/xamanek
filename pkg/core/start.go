package start

import (
	"github.com/cloudflavor/pkg/services"
)

type Config struct{}

type Instance struct {
	Services []*Services
}

func newInstance(config *Config) *Instance {
	return nil
}
