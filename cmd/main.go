package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.Default()
	s := server.NewServer(logger)

	if err := s.Server.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
