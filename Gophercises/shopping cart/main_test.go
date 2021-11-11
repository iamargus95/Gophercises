package main

import "testing"

func TestItemService(t *testing.T) {

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
			offer:  percentOffer{buyQty: 1, percentOff: 50},
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
