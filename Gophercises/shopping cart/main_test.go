package main

import "testing"

func TestItemService(t *testing.T) {

	var testCases = []struct {
		input  item
		output int //shoppingCart.Bill
	}{
		{
			input: item{
				name:     "Dove soap",
				price:    30,
				quantity: 5,
			},

			output: 30 * 5,
		},
		{
			input: item{
				name:          "Dove soap",
				price:         30,
				quantity:      3,
				buyXgetXoffer: buyXgetXoffer{buyQty: 2, getFreeQty: 1},
			},

			output: 30 * 2,
		},
		{
			input: item{
				name:          "Axe Deo",
				price:         100,
				quantity:      2,
				buyXgetXoffer: buyXgetXoffer{},
			},

			output: 100 * 2,
		},
	}

	for _, v := range testCases {
		got := itemService(v.input)
		if got != v.output {
			t.Fail()
		}
	}
}
