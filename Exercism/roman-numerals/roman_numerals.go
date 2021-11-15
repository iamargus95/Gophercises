package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {

	if input <= 0 || input > 3e3 {
		return "", errors.New("out of range for roman numeral")
	}

	result := ""

	convert := []struct {
		value int
		roman string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, res := range convert {
		for input >= res.value {
			result += res.roman
			input -= res.value
		}
	}

	return result, nil
}
