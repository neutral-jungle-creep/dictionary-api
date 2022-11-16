package dto

type WordDto struct {
	Id          int    `json:"id"`
	ForeignWord string `json:"word"`
	Translation string `json:"translation"`
}

func NewWordDto(id int, word string, translation string) *WordDto {
	return &WordDto{
		Id:          id,
		ForeignWord: word,
		Translation: translation,
	}
}
