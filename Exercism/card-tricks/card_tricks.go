package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// a the given index existed in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	if index < len(slice) && index >= 0 {
		return slice[index], true
	}

	return 0, false
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range it is be appended.
func SetItem(slice []int, index, value int) []int {
	if index < len(slice) && index >= 0 {
		slice[index] = value
	} else {
		slice = append(slice, value)
	}
	return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	if length >= 0 {
		newSlice := make([]int, length)
		for i := range newSlice {
			newSlice[i] = value
		}
		return newSlice
	}
	return nil
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < len(slice) && index >= 0 {
		return append(slice[:index], slice[index+1:]...)
	}

	return slice
}
