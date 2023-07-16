package main

import "fmt"

func main(){

var names [4]string
names[0] = "trafalgar"
names[1] = "d"
names[2] = "water"
names[3] = "law"

fmt.Println(names[0], names[1], names[2], names[3])

// Inisialisasi Nilai Awal Array

var fruits = [4]string{"apple", "grape", "banana", "melon"}

fmt.Println("Jumlah element \t\t", len(fruits))
fmt.Println("Isi semua element \t", fruits)

// Inisialisasi Nilai Array Dengan Gaya Vertikal
// var fruits [4]string

// cara horizontal
fruits  = [4]string{"apple", "grape", "banana", "melon"}

// cara vertikal
fruits  = [4]string{
    "apple",
    "grape",
    "banana",
    "melon",
}

// Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
var numbers = [...]int{2, 3, 2, 4, 3}

fmt.Println("data array \t:", numbers)
fmt.Println("jumlah elemen \t:", len(numbers))

// Array Multidimensi
var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

fmt.Println("numbers1", numbers1)
fmt.Println("numbers2", numbers2)

// Perulangan Elemen Array Menggunakan Keyword for
fruits = [4]string{"apple", "grape", "banana", "melon"}

for i := 0; i < len(fruits); i++ {
    fmt.Printf("elemen %d : %s\n", i, fruits[i])
}

// Perulangan Elemen Array Menggunakan Keyword for - range
fruits = [4]string{"apple", "grape", "banana", "melon"}

for i, fruit := range fruits {
    fmt.Printf("elemen %d : %s\n", i, fruit)
}

// Penggunaan Variabel Underscore _ Dalam for - range
fruits = [4]string{"apple", "grape", "banana", "melon"}

for _, fruit := range fruits {
    fmt.Printf("nama buah : %s\n", fruit)
}

// for i, _ := range fruits { }
// atau
// for i := range fruits { }

// Alokasi Elemen Array Menggunakan Keyword make
var fruits2 = make([]string, 2)
fruits2[0] = "apple"
fruits2[1] = "manggo"

fmt.Println(fruits2)  // [apple manggo]

}