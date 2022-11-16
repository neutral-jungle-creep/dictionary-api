package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) showWords(c *gin.Context) {
	words, err := h.service.Show.GetAllWords()
	if err != nil {
		NewException("showWords", c, http.StatusInternalServerError, err).ExceptionResp()
		return
	}
	logrus.Infof("showWords completed successfully, %d", http.StatusOK)
	c.JSON(http.StatusOK, words)
}
