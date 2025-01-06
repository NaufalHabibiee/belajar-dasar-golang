package main

import (
	"fmt"
	"errors"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan NOL")
	} else {
		return nilai / pembagi, nil
	}
}

func Perkalian(nilai int, perkali int) (int, error) {
	if perkali == 0{
		return 0, errors.New("Perkalian dengan NOL")
	}else {
		return nilai * perkali, nil
	}
}

func main() {
	hasil, err := Pembagian(100, 10)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}

	hasilkali, err := Perkalian(10, 10)
	if err == nil {
      fmt.Println("Hasil", hasilkali)
	} else {
		fmt.Println("Error", err.Error())
	}
}

