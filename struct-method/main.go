package main

import "fmt"

type Costumer struct{
	Name string
	Address string
    Age int
}

func (costumer Costumer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", costumer.Name)
	fmt.Println("Hello", name, "my address at", costumer.Address)
	fmt.Println("Hello", name, "my age is", costumer.Age)
}
func main(){

irwan := Costumer {
	Name : "Irwan ganteng",
	Address : "Papua",
	Age : 35,
 }
 fmt.Println(irwan)

hendra := Costumer {
	Name : "Hendra setiawan",
	Address: "Kalimantan",
	Age: 25,
}
fmt.Println(hendra)

irwan.sayHello("ridwan") // irwan sebagai costumer menyapa ridwan
hendra.sayHello("radit") // hendra sebagai costumer menyapa radit


}