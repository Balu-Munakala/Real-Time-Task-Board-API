package server

import (
	"net/http"

	"github.com/Balu-Munakala/Real-Time-Task-Board-API/internal/server/middleware"
	"github.com/Balu-Munakala/Real-Time-Task-Board-API/internal/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RouteHandler *gin.Engine

func NewRouteHandler(logr *zap.Logger, userHandler *user.UserHandler) RouteHandler {
	var router RouteHandler

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestId())
	r.Use(middleware.ZapLogger(logr))

	r.GET("/ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "PONG") })

	r.POST("/register", userHandler.RegisterUser)
	users := r.Group("/users")
	{
		users.GET("/:id", userHandler.GetUserById)
		users.PATCH("/:id", userHandler.EditUserById)
	}

	router = r

	return router
}
