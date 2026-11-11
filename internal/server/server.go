package server

import (
	"net/http"

	"github.com/Balu-Munakala/Real-Time-Task-Board-API/internal/config"
	"github.com/gin-gonic/gin"
	// "go.uber.org/zap"
)

type Server struct {
	HttpServer *http.Server
	// ZapLogger  *zap.Logger
}

// func New(config *config.Config, logger *zap.Logger, router *gin.Engine) Server {
func New(config *config.Config, router *gin.Engine) Server {
	return Server{HttpServer: &http.Server{
		Addr:    config.Addr + ":" + config.Port,
		Handler: router}}
	// },
	// 	ZapLogger: logger,
	// }
}

func (s *Server) Run() {
	s.HttpServer.ListenAndServe()
}
