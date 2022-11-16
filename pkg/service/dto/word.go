package dto

type WordDto struct {
	Id          int    `json:"id"`
	ForeignWord string `json:"word"`
	Translation string `json:"translation"`
}
