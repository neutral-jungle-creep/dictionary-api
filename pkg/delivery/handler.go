package delivery

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	show := router.Group("/show-dictionary")
	{
		show.GET("/words", h.showWords)
	}
	edit := router.Group("/edit-dictionary")
	{
		edit.POST("/edit-word")
		edit.POST("/add-word")
	}
	return router
}

func (h *Handler) showWords(c *gin.Context) {
}
