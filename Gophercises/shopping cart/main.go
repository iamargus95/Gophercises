package main

type item struct {
	name     string
	price    int
	quantity int
	offer    offer
}

type offer interface {
	apply(thing item) int
}

type buyXgetXoffer struct {
	buyQty  int
	freeQty int
}

type percentOffer struct {
	buyQty     int
	percentOff int
}

func main() {

	var o offer
	var i item
	o = buyXgetXoffer{buyQty: 2, freeQty: 1}
	i = item{name: "Dove", price: 30, quantity: 3, offer: o}
}

func (o buyXgetXoffer) apply(thing item) int {

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

	qty := thing.quantity
	deductPercentQty := 0
	itemDiscount := o.buyQty * 2
	for thing.quantity-deductPercentQty >= 0 {
		deductPercentQty += o.buyQty
		thing.quantity -= itemDiscount
	}

	return ((qty - deductPercentQty) * thing.price) + ((deductPercentQty * thing.price) / o.percentOff)
}
