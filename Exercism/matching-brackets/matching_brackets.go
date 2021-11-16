package brackets

func Bracket(input string) bool {
	stack := make([]rune, len(input)+1)
	stackSize := 0
	for _, symbol := range input {
		switch symbol {
		case '(':
			stackSize += 1
			stack[stackSize] = ')'
		case '[':
			stackSize += 1
			stack[stackSize] = ']'
		case '{':
			stackSize += 1
			stack[stackSize] = '}'
		case '}', ']', ')':
			if stackSize == 0 || stack[stackSize] != symbol {
				return false
			}
			stackSize -= 1
		}
	}
	return stackSize == 0
}
