package main

import "fmt"

func getFullName() (string, string) {
	return "Naufal", "Habibie"
}

func main() {
	firstName, lastName := getFullName() // gunakan _ jika ada yang tidak dibutuhkan
	fmt.Println(firstName, lastName)

	// firstName, _ := getFullName()  
	// fmt.Println(firstName)
}