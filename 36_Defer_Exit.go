package main

import "fmt"
import "os"

func main() {
	fmt.Println("\nPenerapan keyword defer\n")
    defer fmt.Println("halo")
    fmt.Println("selamat datang")

	orderSomeFood("pizza")
    orderSomeFood("burger")

	fmt.Println("\nKombinasi defer dan IIFE\n")
	// Agar bisa dimunculkan di akhir blok if, maka harus dibungkus dengan IIFE
	number := 3

    if number == 3 {
        fmt.Println("halo 1")
        func() {
            defer fmt.Println("halo 3")
        }()
    }

    fmt.Println("halo 2")

	fmt.Println("\nPenerapan Fungsi os.Exit()\n")
	// Semua statement setelah exit tidak akan dieksekusi, termasuk juga defer
	defer fmt.Println("halo")
    os.Exit(1)
    fmt.Println("selamat datang")
}

func orderSomeFood(menu string) {
    defer fmt.Println("Terimakasih, silakan tunggu")
    if menu == "pizza" {
        fmt.Print("Pilihan tepat!", " ")
        fmt.Print("Pizza ditempat kami paling enak!", "\n")
        return
    }

    fmt.Println("Pesanan anda:", menu)
}