package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouteHandler *gin.Engine

func NewRouteHandler() RouteHandler {
	var router RouteHandler

	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "PONG") })

	router = r

	return router
}

// func (r *RouteHandler) Use() *gin.Engine {

// }
