package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", homeGetHandler).Methods("GET")
	r.HandleFunc("/", homePostHandler).Methods("POST")
	r.HandleFunc("/register", registerGetHandler).Methods("GET")
	r.HandleFunc("/register", registerPostHandler).Methods("POST")
	fileServer := http.FileServer(http.Dir("./assets"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", fileServer))
	return r
}
