package protein

import "errors"

const (
	// Stop RNA transformation
	Stop = "STOP"
	// Methionine protein
	Methionine = "Methionine"
	// Phenylalanine protein
	Phenylalanine = "Phenylalanine"
	// Leucine protein
	Leucine = "Leucine"
	// Serine protein
	Serine = "Serine"
	// Tyrosine proteinn
	Tyrosine = "Tyrosine"
	// Cysteine protein
	Cysteine = "Cysteine"
	// Tryptophan protein
	Tryptophan = "Tryptophan"
)

var (
	// ErrStop is a flag to stop codon/RNA transformation
	ErrStop error = errors.New("error stop")
	// ErrInvalidBase is an error that returned when the supplied codon is invalid
	ErrInvalidBase error = errors.New("error invalid base")
	// CodonToProtein dictionary
	CodonToProtein = map[string]string{
		"AUG": Methionine,
		"UUU": Phenylalanine,
		"UUC": Phenylalanine,
		"UUA": Leucine,
		"UUG": Leucine,
		"UCU": Serine,
		"UCC": Serine,
		"UCA": Serine,
		"UCG": Serine,
		"UAU": Tyrosine,
		"UAC": Tyrosine,
		"UGU": Cysteine,
		"UGC": Cysteine,
		"UGG": Tryptophan,
		"UAA": Stop,
		"UAG": Stop,
		"UGA": Stop,
	}
)

// FromCodon transform codon to Protein value
func FromCodon(s string) (string, error) {

	v, found := CodonToProtein[s]

	if !found {
		return "", ErrInvalidBase
	}

	if v == Stop {
		return "", ErrStop
	}

	return v, nil
}

// FromRNA transform RNA to protein values
func FromRNA(s string) ([]string, error) {

	var res []string

	codons := splitByN(s, 3)
	for _, codon := range codons {
		protein, err := FromCodon(codon)
		switch err {
		case ErrStop:
			return res, nil
		case ErrInvalidBase:
			return res, err
		default:
			res = append(res, protein)
		}
	}

	return res, nil
}

func splitByN(s string, n int) []string {

	var codons []string

	length := len(s)
	for i := 0; i < length; i++ {
		if i%3 == 0 {
			if i+3 < length {
				codons = append(codons, s[i:i+3])
			} else {
				codons = append(codons, s[i:])
			}
		}
	}

	return codons
}
