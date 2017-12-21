package service

type Service interface {
	StartService() error
	StopService() error
}
