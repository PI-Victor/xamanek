package api

type Node struct {
	IPAddr string
}

type Namespace struct {
	Name      string
	Resources []*Resource
}

type RuntimeEnv struct {
	Name string
	// NOTE: this would probably be a docker image format struct.
	RuntimeImage string
}

type Resource struct {
	Name    string
	Runtime *RuntimeEnv
	// NOTE: Storage should be a docker volume or hostPath
	Storage string
}
