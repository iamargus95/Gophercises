package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {

	if span < 0 {
		return 0, errors.New("span must not be negative")
	}

	if span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	}

	var numbers []int
	for _, r := range digits {
		number, err := strconv.Atoi(string(r))
		if err != nil {
			return 0, errors.New("digits input must only contain digits")
		}
		numbers = append(numbers, number)
	}

	var largestProduct int
	for i := 0; i <= len(numbers)-span; i++ {

		product := 1
		for j := i; j < i+span; j++ {
			product *= numbers[j]
		}

		if product > largestProduct {
			largestProduct = product
		}
	}

	return int64(largestProduct), nil
}
