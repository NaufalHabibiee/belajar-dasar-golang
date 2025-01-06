package main	

import (
	"fmt"
	"belajar-dasar-golang/database"
    _ "belajar-dasar-golang/internal"
)

func main() {
	fmt.Println(database.GetDatabase())

}

