package add

import (
	"errors"
	"log/slog"
	"net/http"

	resp "Order/internal/lib/api/response"
	"Order/internal/lib/logger/sl"
	"Order/internal/lib/random"
	"Order/internal/storage"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	URL   string `json:"url" validate:"required, url"`
	Alias string `json: "alias, omitempty"`
}

type Request struct {
	resp.Response
	Alias string `json: "alias, omitempty"`
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
			slog.String("request_id", middleware.GetReqId(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", sl.Err(err))

			render.JSON(w, r, resp.ValidationError(validateErr))

			return
		}

		alias := req.Alias
		if alias == "" {
			alias = random.NewRandomString(aliasLenght) //what if have generated same alias???
		}

		id, err := adder.AddURL(req.URL, alias)
		if errors.Is(err, storage.ErrUrlExist) {
			log.Info("url already exists", slog.String("url", req.URL))

			render.JSON(w, r, resp.Error("url already exists"))

			return
		}
		if err != nil {
			log.Error("failed to add url", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to add url"))

			return
		}

		log.Info("url added", slog.Int64("id", id))

		responseOK(w, r, alias)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request, alias string) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
		Alias:    alias,
	})
}
