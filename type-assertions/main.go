package main

import "fmt"

func random() any { // any = interface kosong / interface{}
	return "OK"
}

func main() {
	result := random()
	resultString := result.(string)
	fmt.Println(resultString) 
	/* hanya khusus merubah ke tipe data yang 
	   di return oleh interface kosong*/

	// BISA DI GUNAKAN DENGAN SWITCH

	resultRandom := random()
	switch value := resultRandom.(type){
case string :
	fmt.Println("String", value) // jika return dari function random adalah string, maka akan print ini
case int :
	fmt.Println("Int", value) // jika return dari function random adalah integer, maka akan print ini
	default :
	fmt.Println("Uknown", value) // jika function random tidak merenturn dari string dan integer, maka akan print ini
	}
}