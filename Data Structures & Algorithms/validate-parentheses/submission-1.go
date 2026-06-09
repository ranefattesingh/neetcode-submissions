func isValid(s string) bool {
	top := -1
	stack := make([]rune, len(s))

	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			top++
			stack[top] = v
		} else if top > -1 {
			switch v {
				case ')': 
					if stack[top] != '(' {
						return false
					}
				
				case '}':
					if stack[top] != '{' {
						return false
					}
				case ']':
					if stack[top] != '[' {
						return false
					}
			}

			top--
		} else {
			return false
		}
	}

	return top == -1
}
