package time

import "fmt"

func ExampleTime_Calculator(){
	a := fromSeconds(5)
	b := fromHours(600)


	fmt.Println(a.seconds())
	fmt.Println(b.hours())
}