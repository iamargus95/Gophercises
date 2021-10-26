// Package leap has funcs that help determine if an year is a leap year.
package leap

// IsLeapYear checks if input year is a leap year.
func IsLeapYear(year int) bool {
	if year%100 == 0 {
		if year%400 == 0 {
			return true
		} else {
			return false
		}
	}

	if year%4 == 0 {
		return true
	}

	return false
}
