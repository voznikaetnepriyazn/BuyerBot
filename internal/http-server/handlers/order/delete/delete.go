package delete

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

type DeleterURL interface {
	DeleteURL(urlToSave int64) error
}

func New(log *slog.Logger, deleter DeleterURL) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = ""
	}
}
