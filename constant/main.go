package main

import "fmt"

func main() {

	const firstName string = "Muhamad Agung" //tidak bisa diubah
    const lastName = "Naufal Habibie" // tidak bisa diubah

	fmt.Println(firstName)
	fmt.Println(lastName)

}

// INI BISA JUGA
// --------------
// func main_2() {
// 	const (
// 		firstName string = "Budi"
// 		lastName = "Herman"
// 	)

// 	fmt.Println(firstName)
// 	fmt.Println(lastName)
// }