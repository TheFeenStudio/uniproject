package httpHandler

import (
	"fmt"
	"net/http"
	"uniproject/pkg/logger"
)

type HttpHandler struct {
	handler *http.ServeMux
}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (h *HttpHandler) RunService(port int) error {
	h.InitHandler()
	if err := h.StartServer(port); err != nil {
		return err
	}

	return nil
}

func (h *HttpHandler) InitHandler() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(home))

	fileServer := http.FileServer(http.Dir("internal/services/httpHandler/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	h.handler = mux
}

func (h *HttpHandler) StartServer(port int) error {
	logger.Info.Printf("[HttpHandler] server had been started on %d port", port)

	address := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(address, h.handler); err != nil {
		logger.Error.Fatalf("[HttpHandler] error to start server: %s", err)
	}

	return nil
}
