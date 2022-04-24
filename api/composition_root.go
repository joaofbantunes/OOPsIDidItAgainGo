package api

import (
	"github.com/joaofbantunes/OOPsIDidItAgainGo/api/carts"
	"github.com/joaofbantunes/OOPsIDidItAgainGo/app"
	"github.com/joaofbantunes/OOPsIDidItAgainGo/app/command"
	"github.com/joaofbantunes/OOPsIDidItAgainGo/infrastructure"
)

type compositionRoot struct {
	cartRepository          app.CartRepository
	itemRepository          app.ItemRepository
	saleValidatorRepository app.ItemSaleValidatorRepository
}

func NewCompositionRoot() carts.CompositionRoot {
	return compositionRoot{
		cartRepository:          infrastructure.NewInMemoryCartRepository(),
		itemRepository:          infrastructure.NewInMemoryItemRepository(),
		saleValidatorRepository: infrastructure.NewInMemoryItemSaleValidatorRepository(),
	}
}

func (c compositionRoot) CreateAddItemToCartHandler() func(command.AddItemToCart) error {
	return command.NewAddItemToCartHandler(c.cartRepository, c.itemRepository, c.saleValidatorRepository)
}
