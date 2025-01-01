package main

import "fmt"

// BREAK

func main() {
    for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke ", i )
	}

// CONTINUE

for i := 0; i < 10; i++ {
	if i%2 == 0 {
		continue //jika ada perulangan angka genap maka akan di continue
	}
	fmt.Println("Perulangan ke ", i )
}

}