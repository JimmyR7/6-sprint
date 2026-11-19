package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/JimmyR7/6-sprint/internal/server"
)

func main() {
	logger := log.New(os.Stderr, "morse: ", log.LstdFlags)
	srv := server.New(logger)
	if err := srv.HTTP.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Fatal(err)
	}
}
