func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, v := range tokens {
		num, err := strconv.Atoi(v)
		if err == nil {
			stack = append(stack, num)
			continue
		}

		if len(stack) < 2 {
			return -1
		}

		num1 := stack[len(stack)-1]
		num2 := stack[len(stack)-2]

		stack = stack[:len(stack)-2]

		result := 0

		switch v {
			case "+": result = num2 + num1
			case "-": result = num2 - num1
			case "*": result = num2 * num1
			case "/": result = num2 / num1
		}

		stack = append(stack, result)
	}

	return stack[len(stack)-1]
}
