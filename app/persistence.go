package app

import "github.com/joaofbantunes/OOPsIDidItAgainGo/core"

type CartRepository interface {
	GetCartById(id core.CartId) (*core.Cart, error)
	SaveCart(cart *core.Cart) error
}

type ItemRepository interface {
	GetItemById(id core.ItemId) (*core.Item, error)
}

type ItemSaleValidatorRepository interface {
	GetForItem(id core.ItemId) core.ValidateItemSale
}
