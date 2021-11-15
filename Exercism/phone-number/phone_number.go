package phonenumber

import (
	"errors"
	"regexp"
	"strings"
)

//Number format number to american standard phone number
func Number(input string) (string, error) {
	re, _ := regexp.Compile("[0-9]+")
	output := re.FindAllString(input, -1)
	n := strings.Join(output, "")

	if strings.Index(n, "1") == 0 {
		n = n[1:]
	}

	if len(n) != 10 || strings.Index(n, "0") == 0 || strings.Index(n, "1") == 0 || strings.Index(n, "0") == 3 || strings.Index(n, "1") == 3 {
		return "", errors.New("")
	}

	return n, nil
}

//AreaCode return area code of given phone number
func AreaCode(input string) (string, error) {
	output, err := Number(input)
	if err != nil {
		return "", err
	}

	return output[0:3], nil
}

//Format return clean format of given phone number
func Format(input string) (string, error) {
	output, err := Number(input)
	if err != nil {
		return "", err
	}

	return "(" + output[0:3] + ")" + " " + output[3:6] + "-" + output[6:10], nil
}
