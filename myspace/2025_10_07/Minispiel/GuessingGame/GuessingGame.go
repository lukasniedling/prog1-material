package guessingGame

import (
	"fmt"
	"prog1-material/myspace/2025_10_07/Minispiel/NumberGood"
	"prog1-material/myspace/2025_10_07/Minispiel/ReadNumber"
)

func GuessingGame() {
	for n := 0; n < 3; n++ {
		guess := ReadNumber.ReadNumber("Rate eine Zahl: ")
		if NumberGood.NumberGood(guess) {
			fmt.Println("Richtig geraten! :)")
			return
		}
	}
	fmt.Println("Leider falsch! :(")
}
