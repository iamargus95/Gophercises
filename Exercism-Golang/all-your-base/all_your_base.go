package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return []int{}, errors.New("input base must be >= 2")
	}

	if outputBase < 2 {
		return []int{}, errors.New("output base must be >= 2")
	}

	number := 0.0
	for n, in := range inputDigits {
		if in < 0 || in >= inputBase {
			return []int{}, errors.New("all digits must satisfy 0 <= d < input base")
		}
		number += float64(in) * math.Pow(float64(inputBase), float64(len(inputDigits)-n-1))
	}

	outputDigits := []int{}
	if number == 0.0 {
		return []int{0}, nil
	}

	for number > 0 {
		outputDigits = append([]int{int(number) % outputBase}, outputDigits...)
		number = math.Floor(number / float64(outputBase))
	}

	return outputDigits, nil
}
