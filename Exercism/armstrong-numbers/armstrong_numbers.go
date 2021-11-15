package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {

	count := 0
	strN := strconv.Itoa(n)
	power := len(strN)
	for _, val := range strN {
		digit := int(val - '0')
		count = count + int(math.Pow(float64(digit), float64(power)))
	}

	return count == n
}
