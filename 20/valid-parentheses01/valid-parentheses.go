package valid_parentheses01

func isValid(s string) bool {
	m := map[byte]byte{
		'[': ']',
		'{': '}',
		'(': ')',
	}

	stack := []byte{'*'}

	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			stack = append(stack, s[i])
			continue
		}

		index := stack[len(stack)-1]
		if s[i] != m[index] {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 1

}
