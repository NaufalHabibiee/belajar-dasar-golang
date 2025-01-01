package main

import "fmt"

func main() {
	counter := 0

	increment := func() {   // sebuah anonymous function
		fmt.Println("Increment")
		counter ++
	}

	increment()
	increment()
	increment()

	fmt.Println(counter)
}