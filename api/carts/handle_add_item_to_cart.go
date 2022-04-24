package carts

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/joaofbantunes/OOPsIDidItAgainGo/app/command"
	"github.com/joaofbantunes/OOPsIDidItAgainGo/core"
	"github.com/joaofbantunes/OOPsIDidItAgainGo/core/errors"
	"net/http"
)

func registerHandleAddItemToCart(router *mux.Router, handlerFactory func() func(command.AddItemToCart) error) {
	router.HandleFunc(
		"/{cartId}/items",
		func(w http.ResponseWriter, r *http.Request) {
			handleAddItemToCart(w, r, handlerFactory)
		}).Methods(http.MethodPost)
}

func handleAddItemToCart(w http.ResponseWriter, r *http.Request, handlerFactory func() func(command.AddItemToCart) error) {
	vars := mux.Vars(r)
	cartId, parseIdErr := core.ParseCartId(vars["cartId"])

	if parseIdErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c := command.AddItemToCart{CartId: cartId}
	json.NewDecoder(r.Body).Decode(&c)

	commandErr := handlerFactory()(c)
	if commandErr != nil {
		detail := struct {
			Error string `json:"error"`
		}{
			Error: commandErr.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(detail)

		switch commandErr.(type) {
		case errors.Invalid:
			w.WriteHeader(400)
		case errors.NotFound:
			w.WriteHeader(404)
		case errors.Domain:
			w.WriteHeader(409)
		default:
			w.WriteHeader(500)
		}

		return
	}

	w.WriteHeader(204)
}
