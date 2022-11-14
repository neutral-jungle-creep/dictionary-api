package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) showWords(c *gin.Context) {
	c.JSON(http.StatusOK, "hello")
}
