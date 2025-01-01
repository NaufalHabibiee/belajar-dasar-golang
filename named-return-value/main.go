package main

import "fmt"

func getCompleteName()(firstName, middleName, lastName string) {
	firstName = "Agung"
	middleName = "Naufal"
	lastName = "Habibie"

	return
}

func main() {
 a, b, c := getCompleteName()
 fmt.Println(a,b,c)
}