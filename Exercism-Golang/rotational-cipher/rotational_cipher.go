package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	result := make([]rune, len(plain))
	for i, val := range plain {
		iVal := int(val)
		if iVal >= 65 && iVal <= 90 {
			newChar := (int(val-'A') + shiftKey) % 26
			result[i] = rune(int('A') + newChar)
		} else if iVal >= 97 && iVal <= 122 {
			newChar := (iVal - int('a') + shiftKey) % 26
			result[i] = rune(int('a') + newChar)
		} else {
			result[i] = val
		}
	}
	return string(result)
}
