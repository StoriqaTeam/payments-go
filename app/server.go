package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/storiqateam/payments/app/controller"
	"github.com/storiqateam/payments/app/services"
)

// StartServer starts main app http server
func StartServer() error {
	authService := services.AuthService{}
	controller := controller.Controller{AuthService: &authService}
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/sessions", controller.HandleSessionCreate)
	if err := http.ListenAndServe(":8000", r); err != nil {
		return errors.Wrap(err, "starting http server")
	}
	return nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
