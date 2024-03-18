package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

type Server struct {
	handler handler
}

func New(handler handler) *Server {
	return &Server{
		handler: handler,
	}
}

func (s *Server) Start(port string) {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("GET /", fs)

	http.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	http.HandleFunc("GET /{url}", s.handler.GetShortened)
	http.HandleFunc("POST /", s.handler.CreateShortened)

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}
