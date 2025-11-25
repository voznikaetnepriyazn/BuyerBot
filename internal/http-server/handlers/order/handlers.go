package handlers

import (
	"context"
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

type Crud interface {
	NewAdd(log *slog.Logger, adder storage.OrderService) gin.HandlerFunc
	NewDelete(log *slog.Logger, deleter storage.OrderService) gin.HandlerFunc
	NewGetAll(log *slog.Logger, get storage.OrderService) gin.HandlerFunc
	NewGetById(log *slog.Logger, get storage.OrderService) gin.HandlerFunc
	NewUpdate(log *slog.Logger, update storage.OrderService) gin.HandlerFunc
	NewIsOrderCreated(log *slog.Logger, ord storage.OrderService) gin.HandlerFunc
}

func NewAdd(log *slog.Logger, adder storage.OrderService) gin.HandlerFunc {
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

func NewDelete(log *slog.Logger, deleter storage.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.url.delete.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(c.Request.Context())),
		)

		alias := c.Param("id")
		if alias == "" {
			log.Info("id is empty")

			c.JSON(400, gin.H{
				"error": "invalid request",
			})
			return
		}

		err := deleter.DeleteURL(alias)
		if errors.Is(err, storage.ErrUrlNotFound) {
			log.Info("url not found", "id", alias)

			c.JSON(400, gin.H{
				"error": "not found",
			})
		}

		if err != nil {
			log.Error("failed to delete url", sl.Err(err))

			c.JSON(500, gin.H{
				"error": "internal error",
			})

			return
		}

		log.Info("deleted url", slog.String("deleted", alias))

		responseOK(c, alias)
	}
}

func NewGetAll(log *slog.Logger, get storage.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.url.getById.New"

		log = log.With(
			slog.String("op", op),
			slog.Any("request_id", middleware.GetReqIDSlice([]context.Context{c.Request.Context()})),
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

		ids := c.Param("ids")
		if ids == "" {
			log.Info("ids is empty")

			c.JSON(400, gin.H{
				"error": "invalid request",
			})
			return
		}

		resURL, err := get.GetAllURL()
		if errors.Is(err, storage.ErrUrlNotFound) {
			log.Info("urls not found", "ids", ids)

			c.JSON(400, gin.H{
				"error": "not found",
			})
		}

		if err != nil {
			log.Error("failed to get url", sl.Err(err))

			c.JSON(500, gin.H{
				"error": "internal error",
			})

			return
		}

		log.Info("got urls", slog.Any("urls", resURL))

		c.JSON(201, gin.H{
			"urls": resURL,
		})
	}
}

func NewGetById(log *slog.Logger, get storage.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.url.getById.New"

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

		id := c.Param("id")
		if id == "" {
			log.Info("id is empty")

			c.JSON(400, gin.H{
				"error": "invalid request",
			})
			return
		}

		resURL, err := get.GetByIdURL(id)
		if errors.Is(err, storage.ErrUrlNotFound) {
			log.Info("url not found", "id", id)

			c.JSON(400, gin.H{
				"error": "not found",
			})
		}

		if err != nil {
			log.Error("failed to get url", sl.Err(err))

			c.JSON(500, gin.H{
				"error": "internal error",
			})

			return
		}

		log.Info("got url", slog.String("url", resURL))

		c.JSON(201, gin.H{
			"url": resURL,
		})
	}
}

func NewUpdate(log *slog.Logger, update storage.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = ""
	}
}

func NewIsOrderCreated(log *slog.Logger, ord storage.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = ""
	}
}
