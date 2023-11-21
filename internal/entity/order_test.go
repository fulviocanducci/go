package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErrorIfIdIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "Id is required")
}

func TestIfItGetsAnErrorIfPriceIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "price must be greater than zero")
}

func TestIfItGetsAnErrorIfTaxIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "tax must be greater than zero")
}

func TestFinalPrice(t *testing.T) {
	order := Order{Id: "1", Tax: 100.0, Price: 200.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, order.Id, "1")
	assert.Equal(t, order.Tax, 100.0)
	assert.Equal(t, order.Price, 200.0)
	order.CalculateFinalPrice()
	assert.Equal(t, 300.0, order.FinalPrice)
}
