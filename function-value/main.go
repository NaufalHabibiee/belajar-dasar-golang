package main

import "fmt"

func getGoodBye(name string) string {
	dadah := "Good Bye " + name
	return dadah
}

func main() {
	goodbye := getGoodBye
	fmt.Println(goodbye("Opal"))
}