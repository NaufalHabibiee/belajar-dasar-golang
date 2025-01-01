package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = a + b
	var d = b - a
	var e = b / a
	
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

    var i = 10
	i += 10 // i = i+10
	fmt.Println(i)

	i += 5
	fmt.Println(i)

	var j = 1
	j++ //bertambah satu
	j++ // bertambah satu
	fmt.Println(j)

	j--
	fmt.Println(j)

	j--
	fmt.Println(j)

	
}