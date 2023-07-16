package main

import (
    "fmt"
    "math/rand"
    "time"
	"strings"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))
func main() {
	// Penerapan Fungsi
    var names = []string{"John", "Wick"}
    printMessage("halo", names)

	// Fungsi Dengan Return Value / Nilai Balik
	var randomValue int

    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)
    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)
    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)


	// Penggunaan Keyword return Untuk Menghentikan Proses Dalam Fungsi
	divideNumber(10, 2)
    divideNumber(4, 0)
    divideNumber(8, -4)

}

func printMessage(message string, arr []string) {
    var nameString = strings.Join(arr, " ")
    fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
    var value = randomizer.Int()%(max-min+1) + min
    return value
}

func divideNumber(m, n int) {
    if n == 0 {
        fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
        return
    }

    var res = m / n
    fmt.Printf("%d / %d = %d\n", m, n, res)
}