package main

import	"fmt"

func main() {
  var names [3]string
  names[0] = "Opal"
  names[1] = "Budi"
  names[2] = "Jono"
  
  fmt.Println(names)
  fmt.Println(names[0])
  fmt.Println(names[1])
  fmt.Println(names[2])

  // BISA JUGA DIBUAT SECARA LANGSUNG

  var values = [3]int {
	90,
	80,
	70,
  }

 fmt.Println(values)

 var numbers = [...]int{
	90,
	80,
	70,
 }

 fmt.Println(len(numbers)) // untuk mengetahui jumlah array
 fmt.Println(numbers[1]) // untuk mengetahu variabel pada array ke 1
}