package main

import "fmt"

//Exercise 1
func half(n int) (int, bool) {
	h := n / 2
	b := n%2 == 0
	return h, b
}

//Exercise 2
func halfFuncExpression(n int) (int, bool) {
	f := func(m int) bool {return m%2 == 0}
	return n / 2, f(n)
}

//Exercise 3
func theGreatestNum(vn ...int) int {
	g := vn[0]
	for _, num := range vn {
		if num > g { g = num }
	}
	return g
}

//Exercise 5
func foo(vn ...int) {
	fmt.Println(vn)
}

//Exercise 6

func main() {
}
