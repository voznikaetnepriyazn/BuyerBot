package getall

import (
	"log/slog"

	resp "Order/internal/lib/api/response"

	"github.com/gin-gonic/gin"
)

type Response struct {
	URL   string
	alias string
}

type Request struct {
	resp.Response
	Alias string
}

type GetAllURL interface {
	GetAllURL() ([]int64, error)
}

func New(log *slog.Logger, get GetAllURL) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = ""
	}
}
