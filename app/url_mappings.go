package app

import (
	"net/http"

	"github.com/LeQuXi/card/controller"
)

func mapUrls() {
	router.HandleFunc("/cards", controller.Create).Methods(http.MethodPost)
}
