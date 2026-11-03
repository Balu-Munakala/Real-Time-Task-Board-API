package main

import (
	"log"

	"github.com/Balu-Munakala/Real-Time-Task-Board-API/internal/config"
	"github.com/Balu-Munakala/Real-Time-Task-Board-API/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("%v", err)
	}

	router := server.NewRouteHandler()
	s := server.New(cfg, router)

	s.Run()
}
