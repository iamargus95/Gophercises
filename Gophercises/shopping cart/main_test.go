package main

import "testing"

func TestApply(t *testing.T) {

	var testCases = []struct {
		input  item
		offer  offer
		output int
	}{
		{
			input:  item{price: 30, quantity: 5},
			offer:  noOffer{},
			output: 30 * 5,
		},
		{
			input:  item{price: 30, quantity: 3},
			offer:  buyXgetYoffer{buyQty: 2, freeQty: 1},
			output: 30 * 2,
		},
		{
			input:  item{price: 30, quantity: 2},
			offer:  buyXgetYpercentOff{buyQty: 1, percentOff: 50},
			output: 45,
		},
	}

	for _, v := range testCases {
		got := v.offer.apply(v.input)
		if got != v.output {
			t.Fail()
		}
	}
}

func TestGetShoppingCart(t *testing.T) {

	var testCases = []struct {
		cart      shoppingCart
		billOffer billOffer
		billvalue int
	}{
		{
			cart:      shoppingCart{items: []item{{price: 30, quantity: 5, offer: buyXgetYoffer{buyQty: 2, freeQty: 1}}, {price: 100, quantity: 4, offer: buyXgetYpercentOff{buyQty: 1, percentOff: 50}}}},
			billOffer: billOffer{minBillValue: 500, percentOff: 20},
			billvalue: 420,
		},
		{
			cart:      shoppingCart{items: []item{{price: 30, quantity: 10, offer: buyXgetYoffer{buyQty: 2, freeQty: 1}}, {price: 100, quantity: 6, offer: buyXgetYpercentOff{buyQty: 1, percentOff: 50}}}},
			billOffer: billOffer{minBillValue: 500, percentOff: 20},
			billvalue: 504,
		},
	}

	for _, v := range testCases {
		got := v.billOffer.getBill(v.cart)
		if got != v.billvalue {
			t.Fail()
		}
	}
}
