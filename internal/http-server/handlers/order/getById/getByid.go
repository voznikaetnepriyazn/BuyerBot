package getbyid

import (
	"errors"
	"log/slog"

	"Order/internal/http-server/middleware"
	"Order/internal/lib/logger/sl"
	"Order/internal/storage"

	"github.com/gin-gonic/gin"
)

type GetterByIdURL interface {
	GetByIdURL(id string) (string, error)
}

func New(log *slog.Logger, get GetterByIdURL) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.url.getById.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(c.Request.Context())),
		)

		id := c.Param("id")
		if id == "" {
			log.Info("id is empty")

			c.JSON(400, gin.H{
				"error": "invalid request",
			})
			return
		}

		resURL, err := get.GetByIdURL(id)
		if errors.Is(err, storage.ErrUrlNotFound) {
			log.Info("url not found", "id", id)

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

		log.Info("got url", slog.String("url", resURL))

		c.JSON(201, gin.H{
			"url": resURL,
		})
	}
}
