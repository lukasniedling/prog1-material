package PrintMatrix

import "fmt"

func PrintMatrix(arr [][]int) {
	for _, row := range arr {
		fmt.Println(row)
	}
}

func ExamplePrintMatrix() {
	arr := [][]int{
		{1, 2},
		{4, 5},
	}
	PrintMatrix(arr)
	// Output:
	//  [1 2]
	//  [3 4]
}