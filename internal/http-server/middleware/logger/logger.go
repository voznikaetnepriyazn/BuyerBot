package logger

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func New(log *slog.Logger) gin.HandlerFunc {
	log = log.With(
		slog.String("component", "middleware/logger"),
	)

	log.Info("logger middleware enabled")

	fn := func(c *gin.Context) {
		entry := log.With(
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.String("remote_addr", c.ClientIP()),
			slog.String("user_agent", c.GetHeader("UserAgent")),
		)

		t1 := time.Now()
		defer func() {
			entry.Info("request complited",
				slog.Int("status", c.Writer.Status()),
				slog.Int("bytes", c.Writer.Size()),
				slog.String("duration", time.Since(t1).String()),
			)
		}()

		c.Next()
	}
	return fn
}
