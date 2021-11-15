package strand

import (
	"strings"
)

func ToRNA(dna string) string {

	var ret string
	rna := strings.Split(dna, "")

	for i, v := range rna {
		if string(v) == "G" {
			rna[i] = "C"
		}
		if string(v) == "C" {
			rna[i] = "G"
		}
		if string(v) == "T" {
			rna[i] = "A"
		}
		if string(v) == "A" {
			rna[i] = "U"
		}
	}

	ret = strings.Join(rna, "")
	return ret
}
