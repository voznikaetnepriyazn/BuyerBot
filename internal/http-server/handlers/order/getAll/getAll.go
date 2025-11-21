package getall

import (
	"context"
	"errors"
	"log/slog"

	"Order/internal/http-server/middleware"
	"Order/internal/lib/logger/sl"
	"Order/internal/storage"

	"github.com/gin-gonic/gin"
)

type GetAllURL interface {
	GetAllURL() ([]string, error)
}

func New(log *slog.Logger, get GetAllURL) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.url.getById.New"

		log = log.With(
			slog.String("op", op),
			slog.Any("request_id", middleware.GetReqIDSlice([]context.Context{c.Request.Context()})),
		)

		ids := c.Param("ids")
		if ids == "" {
			log.Info("ids is empty")

			c.JSON(400, gin.H{
				"error": "invalid request",
			})
			return
		}

		resURL, err := get.GetAllURL()
		if errors.Is(err, storage.ErrUrlNotFound) {
			log.Info("urls not found", "ids", ids)

			c.JSON(400, gin.H{
				"error": "not found",
			})
		}

		if err != nil {
			log.Error("failed to get url", sl.Err(err))

			c.JSON(500, gin.H{
				"error": "internal error",
			})

			return
		}

		log.Info("got urls", slog.Any("urls", resURL))

		c.JSON(201, gin.H{
			"urls": resURL,
		})
	}
}
