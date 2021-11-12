package main

import "fmt"

type item struct {
	price    int
	quantity int
	offer    offer
}

type offer interface {
	apply(thing item) int
}

type noOffer struct{}

type buyXgetYoffer struct {
	buyQty  int
	freeQty int
}

type buyXgetYpercentOff struct {
	buyQty     int
	percentOff int
}

type shoppingCart struct {
	items []item
	bill  billValue
}

type billValue interface {
	getBill(shoppingCart) int
}

type billOffer struct {
	minBillValue int
	percentOff   int
}

func main() {

	var i item = item{price: 30, quantity: 4, offer: buyXgetYoffer{buyQty: 2, freeQty: 1}}
	fmt.Println(i.offer.apply(i)) //result : 90

	var s shoppingCart
	var b billValue
	s.items = []item{{price: 30, quantity: 10, offer: buyXgetYoffer{buyQty: 2, freeQty: 1}}, {price: 100, quantity: 6, offer: buyXgetYpercentOff{buyQty: 1, percentOff: 50}}}
	fmt.Println(b.getBill(s))
}

func (o buyXgetYoffer) apply(thing item) int {

	qty := thing.quantity
	deductItemQty := 0
	discount := thing.quantity % (o.buyQty + o.freeQty)

	for thing.quantity-discount > o.buyQty {
		deductItemQty += o.freeQty
		thing.quantity -= deductItemQty
	}

	return (qty - deductItemQty) * thing.price
}

func (o buyXgetYpercentOff) apply(thing item) int {

	price := 0
	for i := thing.quantity; i >= 1; i-- {
		if i%(o.buyQty+1) == 0 {
			price += ((thing.price * o.percentOff) / 100)
		} else {
			price += thing.price
		}
	}

	return price
}

func (o noOffer) apply(thing item) int {
	return thing.price * thing.quantity
}

func (b billOffer) getBill(s shoppingCart) int {

	bill := 0
	for _, v := range s.items {
		bill += v.offer.apply(v)
	}

	if bill >= b.minBillValue {
		return (bill - (bill*b.percentOff)/100)
	}

	return bill
}
