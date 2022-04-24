package carts

import (
	"github.com/gorilla/mux"
)

func RegisterCartRoutes(router *mux.Router, compositionRoot CompositionRoot) {
	registerHandleAddItemToCart(router.PathPrefix("/carts").Subrouter(), compositionRoot.CreateAddItemToCartHandler)
}
