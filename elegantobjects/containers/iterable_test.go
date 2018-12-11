package containers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	products := make(Products, 0)
	p1 := products.Create(100)
	assert.NotNil(t, p1)
	assert.NotEmpty(t, products) // fail: Should NOT be empty, but was []
}

func TestAppendingProduct(t *testing.T) {
	products := make(Products, 0)
	p1 := NewProduct(100)
	assert.NotNil(t, p1)
	products = append(products, p1)
	assert.NotEmpty(t, products)
}
