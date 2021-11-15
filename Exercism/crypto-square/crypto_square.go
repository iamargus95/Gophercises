package cryptosquare

import (
	"strings"
	"unicode"
)

func Encode(input string) string {
	input = strings.ToLower(input)

	var inputCleaned []rune
	for _, c := range input {

		if (c >= 'a' && c <= 'z') || unicode.IsNumber(c) {

			inputCleaned = append(inputCleaned, c)
		}
	}

	noCols, noRows := 0, 0
	for noCols*noRows < len(inputCleaned) {
		if noCols > noRows {

			noRows++
		} else {

			noCols++
		}
	}

	rectangle := make([][]string, noCols)
	for i := 0; i < noCols; i++ {

		rectangle[i] = make([]string, noRows)
	}

	row, col := 0, 0
	for i := 0; i < len(inputCleaned); i++ {

		rectangle[row][col] = string(inputCleaned[i])
		row++

		if (i+1)%noCols == 0 {

			col++
			row = 0
		}
	}

	var result strings.Builder
	for i := 0; i < noCols; i++ {

		for j := 0; j < noRows; j++ {

			result.WriteString(rectangle[i][j])

			if rectangle[i][j] == "" {

				result.WriteString(" ")
			}
		}

		if i < noCols-1 {
			result.WriteString(" ")
		}
	}

	return result.String()
}
