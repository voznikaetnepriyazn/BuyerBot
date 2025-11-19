package add

import (
	"errors"
	"log/slog"

	"Order/internal/http-server/middleware"
	resp "Order/internal/lib/api/response"
	"Order/internal/lib/logger/sl"
	"Order/internal/lib/random"
	"Order/internal/storage"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	URL   string `json:"url" validate:"required, url"`
	Alias string `json:"alias,omitempty" binding:"omitempty"`
}

type Request struct {
	resp.Response
	URL   string `json:"url" binding:"required"`
	Alias string `json:"alias,omitempty" binding:"omitempty"`
}

const aliasLenght = 10 //may move to config

type AdderURL interface {
	AddURL(urlToSave string, alias string) (int64, error)
}

func New(log *slog.Logger, adder AdderURL) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.url.add.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(c.Request.Context())),
		)

		var req Request

		if err := c.ShouldBindJSON(&req); err != nil {
			log.Error("failed to decode request body", sl.Err(err))

			c.JSON(400, gin.H{
				"error": "failed to decode request",
			})

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", sl.Err(err))

			c.JSON(400, gin.H{
				"error":   "validation failed",
				"details": formatValidationError(validateErr),
			})

			return
		}

		alias := req.Alias
		if alias == "" {
			alias = random.NewRandomString(aliasLenght) //what if have generated same alias???
		}

		id, err := adder.AddURL(req.URL, alias)
		if errors.Is(err, storage.ErrUrlExist) {
			log.Info("url already exists", slog.String("url", req.URL))

			c.JSON(400, gin.H{
				"error": "url already exists",
			})

			return
		}

		if err != nil {
			log.Error("failed to add url", sl.Err(err))

			c.JSON(500, gin.H{
				"error": "failed to add url",
			})

			return
		}

		log.Info("url added", slog.Int64("id", id))

		responseOK(c, alias)
	}
}

func responseOK(c *gin.Context, alias string) {
	c.JSON(200, gin.H{
		"status": "OK",
		"alias":  alias,
	})
}

func formatValidationError(err validator.ValidationErrors) map[string]string {
	errors := make(map[string]string)
	for _, e := range err {
		errors[e.Field()] = e.Error()
	}
	return errors
}
