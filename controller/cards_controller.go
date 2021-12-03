package controller

import (
	"fmt"
	"net/http"

	"github.com/LeQuXi/card/domain/cards"
	"github.com/LeQuXi/card/services"
)

func Create(w http.ResponseWriter, r *http.Request) {

	if err := oauth.AuthenticateRequest(r); err != nil {

		return
	}
	card := cards.Card{
		User: oauth.GetCallerId(r),
	}
	result, err := services.ItemService.Create(card)
	if err != nil {

	}
	fmt.Println(result)

}
func Get(w http.ResponseWriter, r *http.Request) {

}
