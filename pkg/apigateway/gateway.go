package apigateway

import (
	"github.com/cloudflavor/xamanek/pkg/core"
)

type Resource struct {
	Nodes []*core.Node
}

type APIGateway struct {
	Resources []*Resource
}
