package networkorchestrator

type NetworkOrchestrator interface {
	Start() chan bool
	Stop()
}
