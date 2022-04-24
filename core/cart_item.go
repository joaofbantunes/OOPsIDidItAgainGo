package core

import "github.com/joaofbantunes/OOPsIDidItAgainGo/core/errors"

type CartItem struct {
	itemId   ItemId
	quantity int
}

func NewCartItem(itemId ItemId, quantity int) (*CartItem, error) {
	if quantity < 0 {
		return nil, errors.NewInvalid("quantity must be greater than 0")
	}

	return &CartItem{
		itemId:   itemId,
		quantity: quantity,
	}, nil
}
