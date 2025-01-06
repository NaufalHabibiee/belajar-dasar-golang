package main

import "fmt"

type Address struct{
	City string
    Province string
    Country string
}

func main(){

    // PAS BY VALUE

	// address1 := Address{"Depok", "Jawa Barat", "Indonesia"}
	// address2 := address1

    // address2.City = "Bandung" // pengubahan ini tidak memengaruhi atau merubah city yang berada di address1
    
	// fmt.Println(address1)  // tidak ikut berubah
	// fmt.Println(address2) // berubah menjadi bandung

	// PAS BY REFERENCE

	address1 := Address{"Depok", "Jawa Barat", "Indonesia"}
	address2 := &address1

    address2.City = "Bandung" // address1 akan ikut berubah karena ada menggunakan pointer
    
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi bandung
	


}