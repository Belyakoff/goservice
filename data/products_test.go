package data 

import (
	//"bytes"
	"testing"
	"github.com/stretchr/testify/assert"
)

/*
func TestProductMissingNameReturnsErr(t *testing.T) {
	p := Product{
		Price: 1.22,
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestProductMissingPriceReturnsErr(t *testing.T) {
	p := Product{
		Name:  "abc",
		Price: -1,
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
} */

func TestValidProductDoesNOTReturnsErr(t *testing.T) {
	p := Product{
		Name:  "abc",
		Price: 1.22,
		SKU:   "abc-efg-hji",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err,0)
}