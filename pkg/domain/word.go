package domain

type Word struct {
	Id          int
	ForeignWord string
	Translation string
	Learned     bool
}

func NewWord(id int, word string, translation string, learned bool) *Word {
	return &Word{
		Id:          id,
		ForeignWord: word,
		Translation: translation,
		Learned:     learned,
	}
}
