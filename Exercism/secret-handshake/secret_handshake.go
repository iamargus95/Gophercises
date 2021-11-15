package secret

// Handshake this is the secret code
func Handshake(n uint) []string {
	result := []string{}
	if n&(1<<0) > 0 {
		result = append(result, "wink")
	}

	if n&(1<<1) > 0 {
		result = append(result, "double blink")
	}

	if n&(1<<2) > 0 {
		result = append(result, "close your eyes")
	}

	if n&(1<<3) > 0 {
		result = append(result, "jump")
	}

	if n&(1<<4) > 0 {
		for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}
	}

	return result
}
