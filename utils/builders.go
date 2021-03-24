package utils

import (
	"github.com/gorilla/mux"

)

func BuildBookResource(router *mux.Router, prefix string) {
	//router.HandleFunc(prefix+"/grab", handlers.GetData).Methods("POST")
	//router.HandleFunc(prefix+"/solve", handlers.PushResult).Methods("GET")
}

func BuildManyBooksResourcePrefix(router *mux.Router, prefix string) {
	//router.HandleFunc(prefix, handlers.GetAllBooks).Methods("GET")
}