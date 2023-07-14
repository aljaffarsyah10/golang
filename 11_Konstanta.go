package main
import "fmt"

func main(){

	const firstName string = "john"
fmt.Print("halo ", firstName, "!\n")

const lastName = "wick"
fmt.Print("nice to meet you ", lastName, "!\n")


// Penggunaan Fungsi fmt.Print()
fmt.Println("john wick")
fmt.Println("john", "wick")

fmt.Print("john wick\n")
fmt.Print("john ", "wick\n")
fmt.Print("john", " ", "wick\n")

// Deklarasi Multi Konstanta
const (
    square          = "kotak"
    isToday bool    = true
    numeric uint8   = 1
    floatNum        = 2.2
)

const (
    a = "konstanta"
    b
)

const (
    today string = "senin"
    sekarang
    isToday2 = true
)

const satu, dua = 1, 2
const three, four string = "tiga", "empat"

}
