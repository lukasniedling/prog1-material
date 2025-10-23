package searching

import "fmt"

func ExampleBinFind() {
	l1 := []int{1, 3, 5, 7, 9, 11, 13, 15}
	l2 := []int{}
	
	fmt.Println(BinFind(l1, 1))  // first element
	fmt.Println(BinFind(l1, 15)) // last element
	fmt.Println(BinFind(l1, 7))  // middle element
	fmt.Println(BinFind(l1, 4))  // non-existent element
	fmt.Println(BinFind(l2, 1))  // empty slice

	// Output:
	// 0
	// 7
	// 3
	// -1
	// 
}
