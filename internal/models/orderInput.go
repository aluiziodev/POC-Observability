package models

import (
	"fmt"
	"strings"
)

type CreateOrderInput struct {
	Item     string
	Quantity int
}

func (input *CreateOrderInput) format() {
	input.Item = strings.TrimSpace(input.Item)

}

func (input *CreateOrderInput) Validate() error {
	input.format()
	if strings.TrimSpace(input.Item) == "" {
		return fmt.Errorf("item is required")
	}
	if input.Quantity <= 0 {
		return fmt.Errorf("quantity must be a positive integer")
	}
	return nil
}
