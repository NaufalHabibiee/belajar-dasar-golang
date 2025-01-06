package main

import (
	"belajar-dasar-golang/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Opal")
	fmt.Println(result)

	fmt.Println(helper.SayHello("Budi"))
	// fmt.Println(helper.sayGoodBye("Rahmat")) tidak bisa diakses
	// fmt.Println(helper.version) tidak bisa diakses
	fmt.Println(helper.Application) 
}