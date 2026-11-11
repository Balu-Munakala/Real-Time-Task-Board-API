package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ZapLogger(logr *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		userId, ok := ctx.Get("user_id")
		if !ok {
			userId = "unknown"
		}

		logr.Info("request",
			zap.String("request_id", ctx.MustGet("request_id").(string)),
			zap.String("method", ctx.Request.Method),
			zap.String("path", ctx.Request.URL.Path),
			zap.String("status", strconv.Itoa(ctx.Writer.Status())),
			zap.Duration("latency", time.Since(start)),
			zap.String("client_ip", ctx.ClientIP()),
			zap.String("user_id", userId.(string)),
			// zap.String("errors", ),
			// zap.String("user_id", ),
			// zap.String("trace_id", ),
		)
	}
}
