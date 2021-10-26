package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	histogram := make(Histogram)
	histogram['A'] = 0
	histogram['C'] = 0
	histogram['G'] = 0
	histogram['T'] = 0
	for _, nucleotide := range string(d) {
		if !isValidNucleotide(nucleotide) {
			return nil, errors.New("invalid nucleotide encountered")
		}
		histogram[nucleotide]++
	}
	return histogram, nil
}
func isValidNucleotide(nucleotide rune) bool {
	return nucleotide == 'A' || nucleotide == 'C' || nucleotide == 'G' || nucleotide == 'T'
}
