package isordercreated

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

type OrderCreater interface {
	IsOrderCreated(id int64) (bool, error)
}

func New(log *slog.Logger, ord OrderCreater) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = ""
	}
}
