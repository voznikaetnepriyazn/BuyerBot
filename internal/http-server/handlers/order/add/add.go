package add

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

type AdderURL interface {
	AddURL(urlToSave string, alias string) (int64, error)
}

func New(log *slog.Logger, adder AdderURL) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = ""
	}
}
