package main

import "fmt"

/*
Given a string s which represents an expression, evaluate this expression and return its value.

The integer division should truncate toward zero.
*/

func main() {
	fmt.Println(calculate("3+5 / 2 "))
	fmt.Println(calculate("3+2*2"))
	fmt.Println(calculate("3+(2*2)-(7+5/4)"))
	fmt.Println(calculate("1-1+1"))
}

func calculate(s string) int {

	numStack := make([]int64, 0)
	operStack := make([]byte, 0)

	for i := 0; i < len(s); i++ {

		switch s[i] {
		case ' ':
		case '(', '*', '/', '+', '-':
			operStack = append(operStack, s[i])
		case ')':
			for operStack[len(operStack)-1] != '(' {
				if operStack[len(operStack)-1] == '+' {
					numStack[len(numStack)-2] = numStack[len(numStack)-2] + numStack[len(numStack)-1]
				} else {
					numStack[len(numStack)-2] = numStack[len(numStack)-2] - numStack[len(numStack)-1]
				}
				numStack = numStack[:len(numStack)-1]
				operStack = operStack[:len(operStack)-1]
			}
			operStack = operStack[:len(operStack)-1]
		default:

			j := i
			var num int64
			for ; j < len(s); j++ {
				if !isNum(s[j]) {
					break
				}
				num = num*10 + int64(s[j]-'0')
			}
			i = j - 1

			if len(operStack) == 0 || operStack[len(operStack)-1] == '(' {
				numStack = append(numStack, num)
				break
			}

			switch operStack[len(operStack)-1] {
			case '*':
				numStack[len(numStack)-1] = numStack[len(numStack)-1] * num
				operStack = operStack[:len(operStack)-1]
			case '/':
				numStack[len(numStack)-1] = numStack[len(numStack)-1] / num
				operStack = operStack[:len(operStack)-1]
			default:
				numStack = append(numStack, num)
			}
		}
	}

	for i := range operStack {

		num := numStack[i+1]
		if operStack[i] == '-' {
			num = (-1) * num
		}

		numStack[0] = numStack[0] + num
	}

	return int(numStack[0])
}

func isNum(b byte) bool {
	return int(b-'0') >= 0 && int(b-'0') <= 9
}
