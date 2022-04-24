package core

import (
	"github.com/google/uuid"
	"github.com/joaofbantunes/OOPsIDidItAgainGo/core/errors"
)

type CartId struct {
	value uuid.UUID
}

func NewCartId() CartId {
	return CartId{value: uuid.New()}
}

func ParseCartId(string string) (CartId, error) {
	var cartId CartId
	id, err := uuid.Parse(string)

	if err != nil {
		return cartId, err
	}

	return CartId{value: id}, nil
}

func (id CartId) String() string {
	return id.value.String()
}

type Cart struct {
	id    CartId
	items map[ItemId]*CartItem
}

func NewCart() Cart {
	return Cart{
		id:    NewCartId(),
		items: make(map[ItemId]*CartItem),
	}
}

func NewCartFrom(id CartId, items map[ItemId]*CartItem) Cart {
	return Cart{
		id:    id,
		items: items,
	}
}

func (cart Cart) Id() CartId {
	return cart.id
}

func (cart Cart) AddItemToCart(item Item, quantity int) (*CartItem, error) {

	if _, exists := cart.items[item.id]; exists {
		return nil, errors.NewDomain("item already in cart")
	}

	cartItem, err := NewCartItem(item.id, quantity)

	if err != nil {
		return nil, err
	}

	cart.items[item.id] = cartItem

	return cartItem, nil
}

func (cart Cart) UpdateItemInCart(itemId ItemId, quantity int) (*CartItem, error) {

	if _, exists := cart.items[itemId]; !exists {
		return nil, errors.NewDomain("item not in cart")
	}

	cartItem, err := NewCartItem(itemId, quantity)

	if err != nil {
		return nil, err
	}

	cart.items[itemId] = cartItem

	return cartItem, nil
}

func (cart Cart) RemoveItemFromCart(itemId ItemId) {
	delete(cart.items, itemId)
}
