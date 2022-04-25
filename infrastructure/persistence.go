package infrastructure

import (
	"fmt"
	"github.com/joaofbantunes/OOPsIDidItAgainGo/core"
	"github.com/joaofbantunes/OOPsIDidItAgainGo/core/errors"
)

type InMemoryCartRepository struct {
	carts map[core.CartId]*core.Cart
}

func NewInMemoryCartRepository() InMemoryCartRepository {
	cartId, err := core.ParseCartId("a24ba38b-8421-4ac1-8b48-688d2583a972")
	if err != nil {
		panic("Something strange is going on")
	}

	return InMemoryCartRepository{
		carts: map[core.CartId]*core.Cart{
			cartId: core.NewCartFrom(cartId, make(map[core.ItemId]*core.CartItem)),
		},
	}
}

func (r InMemoryCartRepository) GetCartById(id core.CartId) (*core.Cart, error) {
	if cart, exists := r.carts[id]; exists {
		return cart, nil
	}

	return nil, errors.NewNotFound(fmt.Sprintf("Cart with id \"%s\" not found", id))
}

func (r InMemoryCartRepository) SaveCart(cart *core.Cart) error {
	r.carts[cart.Id()] = cart
	return nil
}

type InMemoryItemRepository struct {
	items map[core.ItemId]*core.Item
}

func NewInMemoryItemRepository() InMemoryItemRepository {
	itemId, err := core.ParseItemId("0b612841-b212-4409-a7aa-f9dd2bdf253e")
	if err != nil {
		panic("Something strange is going on")
	}

	return InMemoryItemRepository{
		items: map[core.ItemId]*core.Item{
			itemId: core.NewItemFrom(itemId),
		},
	}
}

func (r InMemoryItemRepository) GetItemById(id core.ItemId) (*core.Item, error) {
	if item, exists := r.items[id]; exists {
		return item, nil
	}

	return nil, errors.NewNotFound(fmt.Sprintf("Item with id \"%s\" not found", id))
}

type InMemoryItemSaleValidatorRepository struct {
	validators map[core.ItemId]core.ValidateItemSale
	noop       core.ValidateItemSale
}

func NewInMemoryItemSaleValidatorRepository() InMemoryItemSaleValidatorRepository {
	itemId, err := core.ParseItemId("0b612841-b212-4409-a7aa-f9dd2bdf253e")
	if err != nil {
		panic("Something strange is going on")
	}

	return InMemoryItemSaleValidatorRepository{
		validators: map[core.ItemId]core.ValidateItemSale{
			itemId: core.NewComposite(
				core.NewValidateMaximumQuantity(5),
				core.NewValidateMinimumTimeOfDay(core.NewTimeOfDay(12, 0)),
			),
		},
		noop: core.NewNoop(),
	}
}

func (r InMemoryItemSaleValidatorRepository) GetForItem(itemId core.ItemId) core.ValidateItemSale {
	if item, exists := r.validators[itemId]; exists {
		return item
	}

	return r.noop
}
