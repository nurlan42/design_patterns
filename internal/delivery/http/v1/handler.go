package v1

import (
	"github.com/gin-gonic/gin"
	"todo/internal/usecase"

	slog "github.com/sirupsen/logrus"
)

type Handler struct {
	usecase *usecase.Usecase
	log     *slog.Logger
}

// NewHandler takes usecase
func NewHandler(u *usecase.Usecase) *Handler {
	return &Handler{
		usecase: u,
		log:     slog.New(),
	}
}

func (h *Handler) Init(gr *gin.RouterGroup) {
	v1 := gr.Group("/v1")
	{
		h.initTODO(v1)
	}
}
