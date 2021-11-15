package reverse

func Reverse(input string) (reversed string) {

	for _, v := range input {
		reversed = string(v) + reversed
	}
	return
}
