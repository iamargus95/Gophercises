package main

import "fmt"

type item struct {
	name          string
	price         int
	quantity      int
	buyXgetXoffer buyXgetXoffer
	percentOffer  percentOffer
}

type buyXgetXoffer struct {
	buyQty     int
	getFreeQty int
}

type percentOffer struct {
	buyQty     int
	percentOff int
}

func main() {
	fmt.Println(itemService(item{name: "Dove Soap", price: 30, quantity: 5, buyXgetXoffer: buyXgetXoffer{buyQty: 2, getFreeQty: 1}}))
	fmt.Println(itemService(item{name: "Dove Soap", price: 30, quantity: 5, percentOffer: percentOffer{1, 50}}))
}

func itemService(thing item) int {
	if thing.buyXgetXoffer.buyQty == 0 || thing.buyXgetXoffer.getFreeQty == 0 {
		return thing.price * thing.quantity
	}

	qty := thing.quantity
	deductItemQty := 0
	discount := thing.quantity % (thing.buyXgetXoffer.buyQty + thing.buyXgetXoffer.getFreeQty)

	if thing.buyXgetXoffer.buyQty != 0 {
		for thing.quantity-discount > thing.buyXgetXoffer.buyQty {
			deductItemQty += thing.buyXgetXoffer.getFreeQty
			thing.quantity -= deductItemQty
		}

		return (qty - deductItemQty) * thing.price
	}

}
