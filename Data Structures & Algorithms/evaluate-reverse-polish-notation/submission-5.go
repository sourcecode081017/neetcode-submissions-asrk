func evalRPN(tokens []string) int {
	// if its a num, push into stack
	// if its an operator {
	// until stack is empty {
	// num op= stack.pop()
	//}
	// push num back onto stack
	// }
	stack := make([]int, 0)
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err == nil {
			// its a number
			fmt.Println("its a number.. push to stack:", num)
			stack = append(stack, num)
		} else {
			switch token {
				case "+":
					a := stack[len(stack)-1]
					stack = stack[0:len(stack) - 1]
					b := stack[len(stack) - 1]
					stack = stack[0:len(stack) - 1]
					stack = append(stack, a + b)
				case "-":
					fmt.Println("Case -")
					a := stack[len(stack)-1]
					stack = stack[0:len(stack) - 1]
					b := stack[len(stack) - 1]
					stack = stack[0:len(stack) - 1]
					stack = append(stack, b - a)
				case "*":
					fmt.Println("Case *")
					a := stack[len(stack)-1]
					stack = stack[0:len(stack) - 1]
					b := stack[len(stack) - 1]
					stack = stack[0:len(stack) - 1]	
					stack = append(stack, a * b)		
				case "/":
					fmt.Println("Case /")
					a := stack[len(stack)-1]
					stack = stack[0:len(stack) - 1]
					b := stack[len(stack) - 1]
					stack = stack[0:len(stack) - 1]	
					stack = append(stack, b / a)	
			}

		}

	}
	fmt.Println("stack values:", stack)
	return stack[0]
}