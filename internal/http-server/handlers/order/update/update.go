package update

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

type UpdaterURL interface {
	UpdateURL(oldUrl string, urlToSave string, alias string) error
}

func New(log *slog.Logger, update UpdaterURL) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = ""
	}
}
