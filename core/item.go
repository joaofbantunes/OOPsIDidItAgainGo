package core

import (
	"encoding/json"
	"github.com/google/uuid"
)

type ItemId struct {
	value uuid.UUID
}

func NewItemId() ItemId {
	return ItemId{value: uuid.New()}
}

func ParseItemId(string string) (ItemId, error) {
	var itemId ItemId
	id, err := uuid.Parse(string)

	if err != nil {
		return itemId, err
	}

	return ItemId{value: id}, nil
}

func (id ItemId) String() string {
	return id.value.String()
}

func (id *ItemId) UnmarshalJSON(b []byte) error {
	var value string
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}
	parsed, err := uuid.Parse(value)
	if err != nil {
		return err
	}
	id.value = parsed
	return nil
}

type Item struct {
	id ItemId
}

func NewItemFrom(id ItemId) Item {
	return Item{id: id}
}
