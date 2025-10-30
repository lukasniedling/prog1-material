package PrintBoard

import "fmt"

func PrintBoard(board [5][5]string) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Print(cell + " ")
		}
		fmt.Println()
	}
}

func ExamplePrintBoard() {
	board := [5][5]string{
		{" ", "|", " ", "|", " "},
		{"-", "|", "-", "|", "-"},
		{" ", "|", " ", "|", " "},
		{"-", "|", "-", "|", "-"},
		{" ", "|", "", "|", " "},
	}

	PrintBoard(board)
	// Output:
	//	               
	//               
	//               
	//               
	//               
}
