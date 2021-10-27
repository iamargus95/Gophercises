package strain

type Ints []int
type Lists [][]int
type Strings []string

func (r *Ints) Keep(f func(int) bool) (newInts Ints) {

	for _, s := range *r {
		if f(s) {
			newInts = append(newInts, s)
		}
	}
	return newInts
}

func (r *Ints) Discard(f func(int) bool) (newInts Ints) {
	return r.Keep(func(i int) bool {
		return !f(i)
	})
}

func (r *Lists) Keep(f func([]int) bool) Lists {

	newLists := [][]int{}
	for _, s := range *r {
		if f(s) {
			newLists = append(newLists, s)
		}
	}
	return newLists
}

func (r *Strings) Keep(f func(string) bool) Strings {

	newStrings := []string{}
	for _, s := range *r {
		if f(s) {
			newStrings = append(newStrings, s)
		}
	}
	return newStrings

}
