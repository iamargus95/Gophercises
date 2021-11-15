package piglatin

import "strings"

//Sentence translates a string int Pig Latin
func Sentence(s string) string {
	var result string
	pWords := []string{}
	words := strings.Fields(s)
	for _, thisWord := range words {
		pWords = append(pWords, pigify(thisWord))
	}
	result = strings.Join(pWords, " ")
	return result
}

func pigify(s string) string {
	var result string
	switch {
	case s[0] == 'a' || s[0] == 'e' || s[0] == 'i' || s[0] == 'o' || s[0] == 'u':
		result = s + "ay"
	case s[:2] == "yt" || s[:2] == "xr":
		result = s + "ay"
	case s[0] == 'y':
		result = s[1:] + "yay"
	case strings.Contains(s, "qu"):
		i := strings.Index(s, "qu")
		result = s[i+2:] + s[:i+2] + "ay"
	default:
		parts := splitAfterConsonantCluster(s)
		if parts != nil {
			result = parts[1] + parts[0] + "ay"
		}
	}
	return result
}

func splitAfterConsonantCluster(s string) []string {
	result := []string{}
	for i, thisChar := range s {
		if thisChar == 'a' || thisChar == 'e' || thisChar == 'i' || thisChar == 'o' || thisChar == 'u' || thisChar == 'y' {
			result = append(result, s[:i])
			result = append(result, s[i:])
			break
		}
	}
	return result
}
