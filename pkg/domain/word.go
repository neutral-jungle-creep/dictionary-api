package domain

type Word struct {
	Id          int
	Word        string
	Translation string
}

func NewWord(id int, word string, translation string) *Word {
	return &Word{
		Id:          id,
		Word:        word,
		Translation: translation,
	}
}
