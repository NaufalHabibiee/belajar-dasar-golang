package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main(){
	result := getHello("Opal")
    fmt.Println(result)
    
	// JIKA TIDAK MAU SIMPAN KE DALAM VARIABEL

	fmt.Println(getHello("Joko"))
	fmt.Println(getHello("Budi"))
}