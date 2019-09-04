package service

type ServiceStatus int

const(
	Stopped = 1
	Started = 2
	Starting = 3
	Stopping = 4
)

type IService interface{
	Start() error
	Stop() error
	GetStatus() ServiceStatus
}

type Service struct {
	Status ServiceStatus
}

func (t *Service) GetStatus() ServiceStatus {
	return t.Status
}

type ExampleService struct{
	Service
}

func (this *ExampleService) Start() error {
	this.Status = Started
	return nil
}

func (this *ExampleService) Stop() error {
	this.Status = Stopped
	return nil
}