package delivery

import (
	"dictionary/pkg/service/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserRequestWord struct {
	Id          int    `json:"id" binding:"required"`
	ForeignWord string `json:"word" binding:"required"`
	Translation string `json:"translation" binding:"required"`
}

func (h *Handler) addNewWord(c *gin.Context) {
	var input UserRequestWord

	if err := c.BindJSON(&input); err != nil {
		NewException("addNewWord", c, http.StatusBadRequest, err).ExceptionResp()
		return
	}

	wordDto := dto.NewWordDto(input.Id, input.ForeignWord, input.Translation)
	if err := h.service.Edit.AddWord(wordDto); err != nil {
		NewException("addNewWord", c, http.StatusInternalServerError, err).ExceptionResp()
		return
	}

	logrus.Infof("addNewWord completed successfully, %d", http.StatusOK)
	c.JSON(http.StatusOK, "слово успешно добавлено")
}
