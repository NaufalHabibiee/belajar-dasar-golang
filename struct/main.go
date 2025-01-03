package main

import "fmt"

type Costumer struct{
	Name string
	Address string
	Age int
 }

func main() {
	var opal Costumer
    opal.Name = "Muhamad Agung Naufal Habibie"
	opal.Address = "Indonesia"
	opal.Age = 16

	fmt.Println(opal)
	fmt.Println(opal.Name)
	fmt.Println(opal.Address)
	fmt.Println(opal.Age)

	budi := Costumer{
		Name: "Budi Tahalu",
		Address: "Indonesia",
		Age: 20,
	}
	fmt.Println(budi)

	joko:= Costumer{"Joko Kendil", "Indonesia", 30}
    fmt.Println(joko)
}
