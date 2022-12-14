package delivery

import (
	"dictionary/pkg/service"
	"dictionary/pkg/service/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
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
		show.GET("/all-words", h.showWords)
	}
	write := router.Group("/write")
	{
		write.POST("/add-new-word", h.addNewWord)
		write.POST("/edit-word", h.editWord)
	}
	return router
}

func (h *Handler) showWords(c *gin.Context) {
	words, err := h.service.Dictionary.GetAllWords()
	if err != nil {
		NewException("showWords", c, http.StatusInternalServerError, err).ExceptionResp()
		return
	}
	logrus.Infof("showWords completed successfully, %d", http.StatusOK)
	c.JSON(http.StatusOK, words)
}

type AddNewWord struct {
	ForeignWord string `json:"word" binding:"required"`
	Translation string `json:"translation" binding:"required"`
}

func (h *Handler) addNewWord(c *gin.Context) {
	var input AddNewWord

	if err := c.BindJSON(&input); err != nil {
		NewException("addNewWord", c, http.StatusBadRequest, err).ExceptionResp()
		return
	}

	wordDto := dto.NewWordDto(0, input.ForeignWord, input.Translation, false)
	if err := h.service.Dictionary.AddWord(wordDto); err != nil {
		NewException("addNewWord", c, http.StatusInternalServerError, err).ExceptionResp()
		return
	}

	logrus.Infof("addNewWord completed successfully, %d", http.StatusOK)
	c.JSON(http.StatusOK, "слово успешно добавлено")
}

type EditWord struct {
	Id int `json:"id" binding:"required"`
}

func (h *Handler) editWord(c *gin.Context) {
	var input EditWord

	if err := c.BindJSON(&input); err != nil {
		NewException("editWord", c, http.StatusBadRequest, err).ExceptionResp()
		return
	}

	if err := h.service.Dictionary.EditWord(input.Id); err != nil {
		NewException("editWord", c, http.StatusInternalServerError, err).ExceptionResp()
		return
	}

	logrus.Infof("editWord completed successfully, %d", http.StatusOK)
	c.JSON(http.StatusOK, "слово успешно отредактировано")
}
