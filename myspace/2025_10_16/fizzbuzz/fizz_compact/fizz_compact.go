package fizzbuzz

import "fmt"

//Die Funktion geht die Zahlen 1-30 durch und gibt bei Vielfachen von 3 "Fizz", bei Vielfachen von 5 "Buzz" und bei Vielfachen von 3 und 5 "FizzBuzz" aus.
func Fizz(){
	for i := 1; i <= 30; i++ {
		if i%15 == 0 { // Überprüfung ob teilbar durch 3 und 5
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 { // Überprüfung ob teilbar durch 3
			fmt.Println("Fizz")
		} else if i%5 == 0 { // Überprüfung ob teilbar durch 5
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}