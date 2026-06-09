func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, v := range tokens {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			topA := stack[len(stack)-1]
			topB := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch v {
				case "+": stack = append(stack, topB + topA)
				case "-": stack = append(stack, topB - topA)
				case "*": stack = append(stack, topB * topA)
				case "/": stack = append(stack, topB / topA)
			}
		} else {
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}
	}

	return stack[len(stack)-1]
}
