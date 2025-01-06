package main

import "fmt"

// any = interface{} / interface kosong

func Ups() any{
	return 1 /* interface kosong bisa mereturn semua tipe data 
	            seperti string, int, bool */
}

func main(){
	var kosong any = Ups()
	fmt.Println(kosong)
}