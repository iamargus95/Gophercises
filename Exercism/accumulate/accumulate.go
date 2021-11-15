package accumulate

func Accumulate(list []string, transform func(string) string) []string {

	var res []string
	for _, elem := range list {
		res = append(res, transform(elem))
	}
	return res

}
