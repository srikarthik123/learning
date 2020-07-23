package main

import (
	"fmt"
)

func main() {
	fmt.Println("i am just inserted")
	foo()
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}

	}
	lol()
}

func foo() {
	fmt.Println("i am before for loop")
}
func lol() {
	fmt.Println("i am just came out of the loop")
}
