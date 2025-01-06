package main

import "fmt"

func endApp(){
	fmt.Println("End app")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error{
		panic("Ups Error")
	}
}

func main() {
	runApp(true) //jika error false maka akan muncul "End app"
                  //jika error nya true maka akan muncul"Ups Error"
    fmt.Println("Muhamad Agung Naufal Habibie")
}