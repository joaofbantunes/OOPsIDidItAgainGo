package command

import (
	"github.com/joaofbantunes/OOPsIDidItAgainGo/app"
	"github.com/joaofbantunes/OOPsIDidItAgainGo/core"
)

type AddItemToCart struct {
	CartId   core.CartId
	ItemId   core.ItemId
	Quantity int
}

func NewAddItemToCartHandler(
	cartRepository app.CartRepository,
	itemRepository app.ItemRepository,
	saleValidatorRepository app.ItemSaleValidatorRepository) func(AddItemToCart) error {
	return func(command AddItemToCart) error {
		return addItemToCart(cartRepository, itemRepository, saleValidatorRepository, command)
	}
}

func addItemToCart(
	cartRepository app.CartRepository,
	itemRepository app.ItemRepository,
	saleValidatorRepository app.ItemSaleValidatorRepository,
	command AddItemToCart) error {

	cart, getCartErr := cartRepository.GetCartById(command.CartId)
	if getCartErr != nil {
		return getCartErr
	}

	item, getItemErr := itemRepository.GetItemById(command.ItemId)
	if getItemErr != nil {
		return getItemErr
	}

	saleErr := saleValidatorRepository.GetForItem(command.ItemId)(cart, item, command.Quantity)
	if saleErr != nil {
		return saleErr
	}

	_, addItemErr := cart.AddItemToCart(item, command.Quantity)
	if addItemErr != nil {
		return addItemErr
	}

	saveCartErr := cartRepository.SaveCart(cart)
	if saveCartErr != nil {
		return saveCartErr
	}

	return nil
}
