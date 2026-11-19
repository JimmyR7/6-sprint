package server

import (
	"log"
	"net/http"
	"time"

	"github.com/JimmyR7/6-sprint/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	HTTP   *http.Server
}

func New(logger *log.Logger) *Server {
	router := http.NewServeMux()
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/upload", handlers.Upload(logger))
	return &Server{
		Logger: logger,
		HTTP: &http.Server{
			Addr:         ":8080",
			Handler:      router,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}
