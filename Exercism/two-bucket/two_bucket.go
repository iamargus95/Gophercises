package twobucket

import (
	"errors"
)

type bucket struct {
	name  string
	size  int
	level int
}

func (b *bucket) isEmpty() bool {
	return b.level == 0
}
func (b *bucket) isFull() bool {
	return b.level == b.size
}
func (b *bucket) fillUp() {
	b.level = b.size
}
func (b *bucket) emptyOut() {
	b.level = 0
}
func (b *bucket) pourInto(another *bucket) {
	spaceAvailable := another.size - another.level

	var pouredAmount int
	if b.level < spaceAvailable {
		pouredAmount = b.level
	} else {
		pouredAmount = spaceAvailable
	}
	b.level -= pouredAmount
	another.level += pouredAmount
}

var (
	errInvalidParams = errors.New("invalid parameters")
	errNoSolution    = errors.New("no possible solution found")
)

func Solve(sizeBucket1, sizeBucket2, goalAmount int, startBucket string) (goalBucket string, numSteps, otherBucketLevel int, e error) {

	if sizeBucket1 < 1 || sizeBucket2 < 1 || goalAmount < 1 {
		return "", 0, 0, errInvalidParams
	}

	if !relativelyPrime(sizeBucket1, sizeBucket2) {
		return "", 0, 0, errNoSolution
	}
	numSteps = 0
	bucket1 := &bucket{"one", sizeBucket1, 0}
	bucket2 := &bucket{"two", sizeBucket2, 0}

	var this, that *bucket
	switch startBucket {
	case "one":
		this, that = bucket1, bucket2
	case "two":
		this, that = bucket2, bucket1
	default:
		return "", 0, 0, errInvalidParams
	}

	for this.level != goalAmount {
		switch {
		case this.isEmpty():
			this.fillUp()
		case that.isFull():
			that.emptyOut()
		case that.size == goalAmount:
			that.fillUp()
			this, that = that, this
		default:
			this.pourInto(that)
		}
		numSteps++
	}

	return this.name, numSteps, that.level, nil
}

// relativelyPrime returns true, when the only common factor between n1 & n2 is 1.
func relativelyPrime(n1, n2 int) bool {
	for i := 2; i <= n1 && i <= n2; i++ {
		if n1%i == 0 && n2%i == 0 {
			return false
		}
	}
	return true
}
