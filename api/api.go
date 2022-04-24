package api

import (
	"github.com/gorilla/mux"
	"github.com/joaofbantunes/OOPsIDidItAgainGo/api/carts"
	"log"
	"net/http"
)

func StartApi() {
	router := mux.NewRouter().StrictSlash(true)
	carts.RegisterCartRoutes(router, NewCompositionRoot())
	log.Fatal(http.ListenAndServe("localhost:5000", router))
}
