package main

import "fmt"

func main(){
var firstnName string = "john"
var lastName string
lastName = "wick"
//lastName := "wick"
fmt.Printf("halo %s %s!\n", firstnName, lastName)

// var first, second, third string
// first, second, third = "satu", "dua", "tiga"
// var fourth, fifth, sixth string = "empat", "lima", "enam"
// seventh, eight, ninth := "tujuh", "delapan", "sembilan"

// Variabel Underscore _
// Underscore (_) adalah reserved variable yang bisa dimanfaatkan untuk menampung nilai yang tidak dipakai. Bisa dibilang variabel ini merupakan keranjang sampah.
_ = "belajar Golang"
_ = "Golang itu mudah"
// name, _ := "john", "wick"

// Keyword new digunakan untuk membuat variabel pointer dengan tipe data tertentu
// Jika ditampilkan yang muncul bukanlah nilainya melainkan alamat memori nilai tersebut (dalam bentuk notasi heksadesimal)
// name := new(string)
}