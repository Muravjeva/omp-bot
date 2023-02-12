package word

import "github.com/Muravjeva/omp-bot/internal/model/english"

type WordService interface {
	Describe(wordID uint64) (*english.Word, error)
	List(cursor uint64, limit uint64) ([]english.Word, error)
	Create(english.Word) (uint64, error)
	Update(wordID uint64, word english.Word) error
	Remove(wordID uint64) (bool, error)
}

type DummyWordService struct{}

func NewDummyWordService() *DummyWordService {
	return &DummyWordService{}
}
