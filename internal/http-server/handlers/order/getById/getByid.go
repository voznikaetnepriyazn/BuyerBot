package getbyid

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

type GetByIdURL interface {
	GetByIdURL(id int64) (int64, error)
}

func New(log *slog.Logger, get GetByIdURL) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = ""
	}
}
