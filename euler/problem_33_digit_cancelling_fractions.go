package main

import "fmt"

/*
The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to simplify it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.

We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms, find the value of the denominator.
*/

func cancellingFractions(x int, y int) bool {

	x1 := x / 10
	x2 := x % 10
	y1 := y / 10
	y2 := y % 10

	if x2 == 0 || y2 == 0 {
		return false
	}

	if x1 == y1 && x2*y == y2*x ||
		x1 == y2 && x2*y == y1*x ||
		x2 == y1 && x1*y == y2*x ||
		x2 == y2 && x1*y == y1*x {
		return true
	}

	return false
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func prob33() int {

	resultX, resultY := 1, 1

	for i := 11; i < 99; i++ {
		for j := i + 1; j <= 99; j++ {
			if cancellingFractions(i, j) {
				resultX *= i
				resultY *= j
			}
		}
	}

	tmp := gcd(resultX, resultY)
	return resultY / tmp
}

func main() {
	fmt.Println(prob33())
}
