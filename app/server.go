package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

// StartServer starts main app http server
func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8000", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
