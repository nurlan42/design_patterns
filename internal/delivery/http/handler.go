package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurlan42/todo/internal/delivery/http/v1"
	"github.com/nurlan42/todo/internal/usecase"

	swfiles "github.com/swaggo/files"
	ginsw "github.com/swaggo/gin-swagger"
)

type Handler struct {
	Usecase *usecase.Usecase
}

func New(u *usecase.Usecase) *Handler {
	return &Handler{
		Usecase: u,
	}
}

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email martin7.heinz@gmail.com

// @BasePath /api/v1/todo
func (h *Handler) Init() *gin.Engine {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/swagger/*any", ginsw.WrapHandler(swfiles.Handler))

	h.initEndpoints(router)

	return router
}

func (h *Handler) initEndpoints(r *gin.Engine) {
	api := r.Group("/api")

	v1Handler := v1.NewHandler(h.Usecase)

	v1Handler.Init(api)
}
