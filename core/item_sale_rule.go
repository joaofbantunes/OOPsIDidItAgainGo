package core

import (
	"github.com/joaofbantunes/OOPsIDidItAgainGo/core/errors"
	"time"
)

type ValidateItemSale func(cart Cart, item Item, quantity int) error

func validateMaximumQuantity(maximumQuantity int, quantity int) error {
	if quantity > maximumQuantity {
		return errors.NewDomain("Quantity not allowed")
	}

	return nil
}

func NewValidateMaximumQuantity(maximumQuantity int) func(Cart, Item, int) error {
	return func(cart Cart, item Item, quantity int) error {
		return validateMaximumQuantity(maximumQuantity, quantity)
	}
}

func validateMinimumTimeOfDay(minimumTimeOfDay TimeOfDay) error {

	if !IsTimeOfDayReached(time.Now(), minimumTimeOfDay) {
		return errors.NewDomain("Can't buy that yet!")
	}

	return nil
}

func IsTimeOfDayReached(time time.Time, timeOfDay TimeOfDay) bool {
	if time.Hour() > timeOfDay.hour || (time.Hour() == timeOfDay.hour && time.Minute() >= timeOfDay.minute) {
		return true
	}
	return false
}

func NewValidateMinimumTimeOfDay(minimumTimeOfDay TimeOfDay) func(Cart, Item, int) error {
	return func(Cart, Item, int) error {
		return validateMinimumTimeOfDay(minimumTimeOfDay)
	}
}

func NewNoop() func(Cart, Item, int) error {
	return func(Cart, Item, int) error {
		return nil
	}
}

func NewComposite(validators ...ValidateItemSale) func(Cart, Item, int) error {
	return func(cart Cart, item Item, quantity int) error {
		for _, validator := range validators {
			if err := validator(cart, item, quantity); err != nil {
				return err
			}
		}
		return nil
	}
}
