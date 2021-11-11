package main

import "fmt"

type item struct {
	price    int
	quantity int
}

type offer interface {
	apply(thing item) int
}

type noOffer struct{}

type buyXgetYoffer struct {
	buyQty  int
	freeQty int
}

type percentOffer struct {
	buyQty     int
	percentOff int
}

func main() {

	var o offer
	var i item = item{price: 30, quantity: 4}

	o = buyXgetYoffer{buyQty: 2, freeQty: 1}
	fmt.Println(o.apply(i)) //result : 90

	o = percentOffer{buyQty: 1, percentOff: 50}
	fmt.Println(o.apply(i)) //result : 90

	o = noOffer{}
	fmt.Println(o.apply(i)) //result : 120

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

func (o percentOffer) apply(thing item) int {

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
