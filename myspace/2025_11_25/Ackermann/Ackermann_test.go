package ackermann

import "fmt"

func ExampleAckermann() {
	fmt.Println(Ackermann(1, 0))
	fmt.Println(Ackermann(1, 2))
	fmt.Println(Ackermann(2, 2))
	fmt.Println(Ackermann(3, 3))
	// Output:
	// 1
	// 4
	// 7
	// 61
}

func Ackermann(m, n int) int {
	if m == 0 {
		return n + 1
	}
	if n == 0 {
		return Ackermann(m-1, 1)
	}
	return Ackermann(m-1, Ackermann(m, n-1))
}
