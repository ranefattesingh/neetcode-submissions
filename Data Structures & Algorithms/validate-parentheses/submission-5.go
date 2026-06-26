func isValid(s string) bool {
    stack := []rune{}


    for _, v := range s {
        // push when encountered opening {, [, or (.
        if v == '{' || v == '[' || v == '(' {
            stack = append(stack, v)
            continue
        }

        // stack is empty and incoming element is not a pushable.
        if len(stack) == 0 {
            return false
        }

        // match not found for existing opening {, [, or (.
        if !match(v, stack[len(stack)-1]) {
            return false
        }

        // pop the stack top
        stack = stack[:len(stack)-1]
    }

    return len(stack) == 0
}

func match(incoming, top rune) bool {
    return (incoming == '}' && top == '{') || (incoming == ']' && top == '[' || incoming == ')' && top == '(')
}
