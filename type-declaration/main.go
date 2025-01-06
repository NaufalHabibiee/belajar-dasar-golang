package main

import "fmt"

func main() {
	
	type NoKTP string
	

	var ktpOpal NoKTP = "1111111111"
	
	var contoh string = "2222222222"
	var  contohKtp NoKTP = NoKTP(contoh)

fmt.Println(ktpOpal)
fmt.Println(contoh)
fmt.Println(contohKtp)


}