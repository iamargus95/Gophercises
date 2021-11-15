package raindrops

import "strconv"

func Convert(drops int) string {
	if drops%3 != 0 && drops%5 != 0 && drops%7 != 0 {
		return strconv.Itoa(drops)
	}

	var res string
	if drops%3 == 0 {
		res += "Pling"
	}

	if drops%5 == 0 {
		res += "Plang"
	}

	if drops%7 == 0 {
		res += "Plong"
	}

	return res
}
