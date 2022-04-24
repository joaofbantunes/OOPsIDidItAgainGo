package carts

import "github.com/joaofbantunes/OOPsIDidItAgainGo/app/command"

type CompositionRoot interface {
	CreateAddItemToCartHandler() func(command.AddItemToCart) error
}
