package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// StartServer starts main app http server
func StartServer() error {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	if err := http.ListenAndServe(":8000", r); err != nil {
		return errors.Wrap(err, "starting http server")
	}
	return nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
