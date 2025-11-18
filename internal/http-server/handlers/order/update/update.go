package update

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Response struct {
	URL   string
	alias string
}

type Request struct {
	//resp.Response
	Alias string
}

type UpdateURL interface {
	UpdateURL(oldUrl string, urlToSave string, alias string) error
}

func New(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context)
}
