package middleware

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := fmt.Sprintf("%d", time.Now().UnixNano())
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "requestID", id))
		c.Header("X-Request-ID", id)
		c.Next()
	}
}

func Recoverer() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Error("panic recovered", slog.Any("error", err))
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})

			c.Abort()
		}()

		c.Next()
	}
}
