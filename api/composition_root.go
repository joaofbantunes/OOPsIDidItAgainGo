package api

import (
	"fmt"
	"github.com/joaofbantunes/OOPsIDidItAgainGo/api/carts"
	"github.com/joaofbantunes/OOPsIDidItAgainGo/app"
	"github.com/joaofbantunes/OOPsIDidItAgainGo/app/command"
	"github.com/joaofbantunes/OOPsIDidItAgainGo/infrastructure"
	"github.com/sirupsen/logrus"
	"reflect"
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
	return func(cart command.AddItemToCart) error {
		return commandLogger(cart, command.NewAddItemToCartHandler(c.cartRepository, c.itemRepository, c.saleValidatorRepository))
	}
}

func commandLogger[T any](c T, decorated func(T) error) error {
	log := logrus.WithFields(map[string]interface{}{
		"command":     c,
		"commandName": reflect.TypeOf(c),
	})
	log.Info(fmt.Sprintf("Executing command %T", c))
	defer log.Info(fmt.Sprintf("Finished executing command %T", c))
	return decorated(c)
}
