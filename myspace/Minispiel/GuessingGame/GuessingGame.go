package guessingGame
import "fmt"
import "prog1-material/myspace/Minispiel/NumberGood"
import "prog1-material/myspace/Minispiel/ReadNumber"

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

