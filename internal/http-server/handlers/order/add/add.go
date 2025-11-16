package add

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func New(log *slog.Logger, urlSaver URLSaver) gin.HandlerFunc {
	return func(c *gin.Context)
}
