package server

import (
	"net/http"

	"github.com/Balu-Munakala/Real-Time-Task-Board-API/internal/config"
	"github.com/gin-gonic/gin"
	// "github.com/Balu-Munakala/Real-Time-Task-Board-API/internal/platform/logger"
)

type Server struct {
	HttpServer *http.Server
	// UberLogger *uber.Logger
}

func New(config *config.Config, router *gin.Engine) Server {
	// handler.Use(logger)

	return Server{HttpServer: &http.Server{
		Addr:    config.Addr + ":" + config.Port,
		Handler: router,
	}}
}

func (s *Server) Run() {
	s.HttpServer.ListenAndServe()
}
