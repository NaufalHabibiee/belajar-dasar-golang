package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string){
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string{
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Opal", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

	

}

// Bisa juga menggunakan TYPE

type Filter func(string) string
    
func sayHelloWithFilterr(name string, filter Filter){
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}