package perfect

import "errors"

// Define the Classification type here.
type Classification int

var (
	ClassificationPerfect   Classification = 3
	ClassificationAbundant  Classification = 2
	ClassificationDeficient Classification = 1
	ErrOnlyPositive         error          = errors.New("only positives")
)

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	} else if n == 1 {
		return ClassificationDeficient, nil
	}

	count := int64(1)

	for f := int64(2); f < n; f++ {
		if n%f == 0 {
			count += f
		}
	}

	if n == count {
		return ClassificationPerfect, nil
	} else if count > n {
		return ClassificationAbundant, nil
	} else {
		return ClassificationDeficient, nil
	}

}
