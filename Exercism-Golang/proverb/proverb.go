package proverb

import "fmt"

const final = "And all for the want of a %s."
const piece = "For want of a %s the %s was lost."

func Proverb(rhyme []string) (p []string) {

	if len(rhyme) == 0 {
		return []string{}
	}

	p = make([]string, 0)

	var v string

	for i := range rhyme {
		if i == len(rhyme)-1 {
			v = fmt.Sprintf(final, rhyme[0])
		} else {
			v = fmt.Sprintf(piece, rhyme[i], rhyme[i+1])
		}
		p = append(p, v)
	}
	return p
}
