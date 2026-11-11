package middleware

import (
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func RequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := ctx.GetHeader("X-Request-ID")
		if requestId == "" {
			requestId = uuid.NewString()
		}

		ctx.Header("X-Request-ID", requestId)
		ctx.Set("request_id", requestId)
		ctx.Next()
	}
}
