package main

import "fmt"

func main() {
	name := "Opal"

	if name == "Opal" {
		fmt.Println("Hello Opal")
	} else if  name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	// PENGGUNAAN IF STATEMENT

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}