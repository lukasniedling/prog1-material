package ReadNumber

import "fmt"

func ReadNumber(prompt string) int { 
	var n int
	fmt.Println("Gib eine Zahl ein um zu raten.")
	fmt.Scan(&n)
	return n
}
