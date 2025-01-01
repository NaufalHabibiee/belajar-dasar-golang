package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	fmt.Println("Selesai")

	// FOR DENGAN STATEMENT

	 for numbers := 1; numbers <= 10; numbers++{
		fmt.Println("Perulangan ke ", numbers)
	}

	names := []string{"Agung", "Naufal", "Habibie"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}


}