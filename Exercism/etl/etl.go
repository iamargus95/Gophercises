package etl

import "strings"

func Transform(old map[int][]string) map[string]int {
	newScore := make(map[string]int)
	for score, list := range old {
		for _, letter := range list {
			newScore[strings.ToLower(letter)] = score
		}
	}
	return newScore
}
