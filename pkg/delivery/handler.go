package delivery

import (
	"dictionary/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	show := router.Group("/show")
	{
		show.GET("/words", h.showWords)
	}
	return router
}
