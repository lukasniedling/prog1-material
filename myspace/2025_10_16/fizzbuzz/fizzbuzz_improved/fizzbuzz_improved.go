package fizzbuzzimproved

import "fmt"

//Die Funktion geht die Zahlen 1-30 durch und gibt bei Vielfachen von 3 "Fizz", bei Vielfachen von 5 "Buzz" und bei Vielfachen von 3 und 5 "FizzBuzz" aus.
func FizzImproved(x, y, n int){
	for i := 1; i <= n; i++ {
		if i%x == 0 && i%y == 0 { // Überprüfung ob teilbar durch 3 und 5
			fmt.Println("FizzBuzz")
		} else if i%x == 0 { // Überprüfung ob teilbar durch 3
			fmt.Println("Fizz")
		} else if i%y == 0 { // Überprüfung ob teilbar durch 5
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}