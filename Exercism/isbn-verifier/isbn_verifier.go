package isbn

func IsValidISBN(isbn string) bool {
	count := 10
	result := 0
	for _, digit := range isbn {
		switch {
		case digit == '-':
			continue
		case digit >= '0' && digit <= '9':
			result += count * int(digit-'0')
			count--
		case digit == 'X' && count == 1:
			result += 10
			count--
		default:
			return false
		}
	}
	return result%11 == 0 && count == 0
}
