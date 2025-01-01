package main

import "fmt"

func main(){
	name := "Opal"
    
    switch name {
	case "Opal" :
		fmt.Println("Hello Opal")
	case "Joko" :
		fmt.Println("Hello Joko")
	default: 
	fmt.Println("Hi, Boleh Kenalan?")
	}

	// SWITCH SHORT STATEMENT

	switch length := len(name); length > 5{
	case true :
		fmt.Println("Nama terlalu panjang")
	case false :
		fmt.Println("Nama sudah benar")
	}

    // SWITCH TANPA KONDISI

	name = "OpalOpalOpal"
	length := len(name)
	switch{
	case length > 10 :
			fmt.Println("Nama terlalu panjang")
		
	case length > 5 :
			fmt.Println("Nama lumayan panjang")

	default:
		fmt.Println("Nama sudah benar")
		
	}


	

}