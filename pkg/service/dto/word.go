package dto

type WordDto struct {
	Id          int    `json:"id"`
	ForeignWord string `json:"word"`
	Translation string `json:"translation"`
	Learned     bool   `json:"learned"`
}

func NewWordDto(id int, word string, translation string, learned bool) *WordDto {
	return &WordDto{
		Id:          id,
		ForeignWord: word,
		Translation: translation,
		Learned:     learned,
	}
}
