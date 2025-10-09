package ReadNumber

import "fmt"

func ReadNumber(prompt string) int { 
	var n int
	fmt.Println(prompt)
	fmt.Scan(&n)
	return n
}
