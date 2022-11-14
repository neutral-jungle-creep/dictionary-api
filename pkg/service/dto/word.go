package dto

type WordDto struct {
	Id          int    `json:"id" binding:"required"`
	ForeignWord string `json:"word" binding:"required"`
	Translation string `json:"translation" binding:"required"`
}
