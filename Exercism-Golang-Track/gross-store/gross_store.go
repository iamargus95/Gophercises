package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{"quarter_of_a_dozen": 3, "half_of_a_dozen": 6, "dozen": 12, "small_gross": 120, "gross": 144, "great_gross": 1728}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	qty, exists := units[unit]
	if !exists {
		return false
	}
	bill[item] = qty
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	qty, exists := units[unit]
	if !exists {
		return false
	}

	num, exists := bill[item]
	if !exists {
		return false
	}

	if num < qty {
		return false
	}

	if num == qty {
		delete(bill, item)
		return true
	}

	bill[item] -= qty
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	num, exists := bill[item]
	if !exists {
		return 0, false
	}
	return num, true
}
