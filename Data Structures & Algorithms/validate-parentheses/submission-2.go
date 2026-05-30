func isValid(s string) bool {
    stack := make([]rune, 0)
	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
		} else {
			if len(stack) > 0 {
				c := stack[len(stack) - 1]
				stack = stack[:len(stack) - 1]
				switch c {
					case '(':
						if r != ')' {
							return false
						}
					case '[':
						if r != ']' {
							return false
						}
					case '{':
						if r != '}' {
							return false
						}
				}
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
