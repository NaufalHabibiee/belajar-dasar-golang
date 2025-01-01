package main

import "fmt"

func main() {
	names := [...]string{"Opal", "Budi", "Agung", "Marsya", "Ambar", "Herman"}
	slice1 := names[4:6] // dari array 4 sampe 5
	slice2 := names[:3] // berhenti sampai array ke 2
	slice3 := names[3:] // mulai dari array ke 3
	slice4 := names[:] // dari awal sampai akhir
	

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

wall := "------------------------------------"
fmt.Println(wall)

	days := [...]string {"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(days) // [senin, selasa, rabu, kamis, jumat, sabtu baru, minggu baru]

wall2 := "------------------------------------"
fmt.Println(wall2)

    daySlice2 := append(daySlice1, "libur baru") // [sabtu baru, minggu baru, libur baru]
	daySlice2[0] = "sabtu lama"



	fmt.Println(daySlice2) // [sabtu lama, minggu baru, libur baru]
	
wall3 := "------------------------------------"
fmt.Println(wall3)
	
	fmt.Println(days) // senin, selasa, rabu, kamis, jumat, sabtu baru, minggu baru]

newSlice := make([]string, 2, 5) // 2 = panjang slice, 5 = kapasitas slice
newSlice[0] = "Opal"
newSlice[1] = "Opal"
// newSlice[2] = "Opal" // error harusnya menggunakan append

fmt.Println(newSlice)
fmt.Println(len(newSlice)) // untuk mengetahui panjang slice
fmt.Println(cap(newSlice)) // untuk mengetahui kapasitas slice
  
newSlice2 := append(newSlice, "Opal")
fmt.Println(newSlice2)
fmt.Println(len(newSlice2))
fmt.Println(cap(newSlice2))

newSlice2[0] = "Budi"
fmt.Println(newSlice2)
fmt.Println(newSlice)

fromSlice := days[:]
toSlice := make([]string, len(fromSlice), cap(fromSlice))

copy(toSlice, fromSlice)

fmt.Println(fromSlice)
fmt.Println(toSlice)

numbers := [...]int {90, 80, 70, 60}

fromSlice2 := numbers[:]
toSlice2 := make([]int, len(fromSlice2), cap(fromSlice2))

copy(toSlice2, fromSlice2)

fmt.Println(fromSlice2)
fmt.Println(toSlice2)

// PERBEDAAN ARRAY DAN SLICE

iniArray := [...]string {"opal", "budi"}
iniSlice := []string {"opal", "budi"}

fmt.Println(iniArray)
fmt.Println(iniSlice )
}