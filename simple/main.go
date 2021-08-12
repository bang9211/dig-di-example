package main

import (
	"fmt"
	"net/http"

	"go.uber.org/dig"
)

type Handler struct {
	Greeting string
	Path     string
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s from %s", h.Greeting, h.Path)
}

func NewHello1Handler() HandlerResult {
	return HandlerResult{
		Handler: Handler{
			Path:     "/hello1",
			Greeting: "welcome",
		},
	}
}

func NewHello1Handler2() HandlerResult {
	return HandlerResult{
		Handler: Handler{
			Path:     "/hello2",
			Greeting: "welcome2",
		},
	}
}

type HandlerResult struct {
	Handler Handler
}

func RunServer(param HandlerResult) error {
	mux := http.NewServeMux()
	mux.Handle(param.Handler.Path, param.Handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func main() {
	container := dig.New()

	container.Provide(NewHello1Handler)
	container.Provide(NewHello1Handler2)

	container.Invoke(RunServer)
}
