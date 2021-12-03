package services

import (
	"net/http"

	"github.com/LeQuXi/card/domain/cards"
	"github.com/LeQuXi/lang/utils/errors"
)

var (
	CardsService cardsServiceInterface = &cardsService{}
)

type cardsServiceInterface interface {
	Create(cards.Card) (*cards.Card, errors.RestErr)
	Get(string) (*cards.Card, errors.RestErr)
}
type cardsService struct {
}

func (s *cardsService) Create(cards.Card) (*cards.Card, errors.RestErr) {
	return nil, errors.NewInternalServerError("implement me", http.StatusNotImplemented, "not_implemented", nil)
}
func (s *cardsService) Get(string) (*cards.Card, errors.RestErr) {
	return nil, errors.NewInternalServerError("implement me", http.StatusNotImplemented, "not_implemented", nil)
}
