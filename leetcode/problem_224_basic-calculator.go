package main

import "fmt"

/*
Given a string s representing an expression, implement a basic calculator to evaluate it.
*/

func main() {
	fmt.Println(calculate("1 + 1"))
	fmt.Println(calculate(" 2-1 + 2 "))
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
	fmt.Println(calculate("2147483647"))
	fmt.Println(calculate("-2+ 1"))
	fmt.Println(calculate("- (3 + (4 + 5))"))
}

func calculate(s string) int {

	numStack := make([]int64, 0)
	oprStack := make([]byte, 0)

	if s[0] == '+' || s[0] == '-' {
		s = "0" + s
	}

	for i := 0; i < len(s); i++ {
		// if is space, continue
		if s[i] == ' ' {
			continue
		}
		// if is (
		if s[i] == '(' {
			oprStack = append(oprStack, s[i])
			continue
		}
		// if is + / -
		if isOperator(s[i]) {

			if i == 0 || s[i-1] == '(' {

				var num int64
				j := i + 1
				for ; j < len(s); j++ {
					if !isNum(s[j]) {
						break
					}
					num = num*10 + int64(s[j]-'0')
				}

				if s[i] == '-' {
					num = (-1) * num
				}
				numStack = append(numStack, num)

				i = j - 1
				continue
			}

			oprStack = append(oprStack, s[i])
			continue
		}

		if s[i] == ')' {
			// pop (
			oprStack = oprStack[:len(oprStack)-1]

			if len(oprStack) == 0 {
				continue
			}

			if oprStack[len(oprStack)-1] == '+' {
				numStack[len(numStack)-2] = numStack[len(numStack)-2] + numStack[len(numStack)-1]
			} else {
				numStack[len(numStack)-2] = numStack[len(numStack)-2] - numStack[len(numStack)-1]
			}
			oprStack = oprStack[:len(oprStack)-1]
			numStack = numStack[:len(numStack)-1]

			continue
		}
		// if is num
		if isNum(s[i]) {

			var num int64
			j := i
			for ; j < len(s); j++ {
				if !isNum(s[j]) {
					break
				}
				num = num*10 + int64(s[j]-'0')
			}

			i = j - 1

			// if oprStack is empty or is not operator
			if len(oprStack) == 0 || !isOperator(oprStack[len(oprStack)-1]) {
				numStack = append(numStack, num)
				continue
			}

			if oprStack[len(oprStack)-1] == '+' {
				numStack[len(numStack)-1] = numStack[len(numStack)-1] + num
			} else {
				numStack[len(numStack)-1] = numStack[len(numStack)-1] - num
			}
			oprStack = oprStack[:len(oprStack)-1]
		}
	}

	if len(numStack) == 0 {
		return 0
	}
	return int(numStack[0])
}

func isNum(r byte) bool {
	return int(r-'0') >= 0 && int(r-'0') <= 9
}

func isOperator(r byte) bool {
	return r == '+' || r == '-'
}
