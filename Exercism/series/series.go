package series

func All(n int, s string) (res []string) {
	for i := 0; i <= len(s)-n; i++ {
		res = append(res, s[i:i+n])
	}
	return
}

func UnsafeFirst(n int, s string) string {
	res, ok := First(n, s)
	if ok {
		return res
	}
	panic(nil)
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return s, false
	}
	return s[0:n], true
}
