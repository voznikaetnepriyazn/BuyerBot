package isordercreated

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

type IsOrderCreated interface {
	IsOrderCreated(id int64) (bool, error)
}

func New(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context)
}
