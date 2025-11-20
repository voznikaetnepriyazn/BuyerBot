package getall

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type GetAllURL interface {
	GetAllURL() ([]string, error)
}

func New(log *slog.Logger, get GetAllURL) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = ""
	}
}
