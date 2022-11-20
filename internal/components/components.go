package components

type HttpServer interface {
	InitHandler()
	StartServer(port int) error
	RunService(port int) error
}

type App interface {
	Run()
	Setup()
	Stop()
}
