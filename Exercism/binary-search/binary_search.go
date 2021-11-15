package binarysearch

func SearchInts(sorted []int, key int) int {

	lower, upper := 0, len(sorted)-1
	for upper >= lower {
		middle := lower + (upper-lower)/2
		current := sorted[middle]
		switch true {
		case current < key:
			lower = middle + 1
		case current > key:
			upper = middle - 1
		default:
			return middle
		}
	}

	return -1
}
