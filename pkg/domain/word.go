package domain

type Word struct {
	Id          int
	ForeignWord string
	Translation string
}

func NewWord(id int, word string, translation string) *Word {
	return &Word{
		Id:          id,
		ForeignWord: word,
		Translation: translation,
	}
}
