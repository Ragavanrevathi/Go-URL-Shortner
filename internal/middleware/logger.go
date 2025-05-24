package middleware

import (
	"context"
	"net/http"
	"shorten-url/pkg/common"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var traceID string

		if c.Request.Method == http.MethodGet && c.FullPath() == "/:shortKey" {
			traceID = c.Param("shortKey")
			if traceID == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Missing short URL key in path"})
				c.Abort()
				return
			}
		} else {
			traceID = c.GetHeader("traceId")
			if traceID == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Trace-Id header is required"})
				c.Abort()
				return
			}
		}

		// Enrich logger with traceId
		ctxLogger := logger.With(zap.String("traceId", traceID))

		// Create new context with logger and traceId
		ctx := context.WithValue(c.Request.Context(), common.TraceIDKey, traceID)
		ctx = context.WithValue(ctx, common.LoggerKey, ctxLogger)

		// Set updated context back to request
		c.Request = c.Request.WithContext(ctx)

		// Add traceId to response header
		c.Header("traceId", traceID)

		c.Next()
	}
}
