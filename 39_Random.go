package main

import (
    "fmt"
    "math/rand"
	"time"
)

var randomizer3 = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	fmt.Println("\nPackage math/rand\n")
    randomizer := rand.New(rand.NewSource(10))
    fmt.Println("random ke-1:", randomizer.Int()) // 5221277731205826435
    fmt.Println("random ke-2:", randomizer.Int()) // 3852159813000522384
    fmt.Println("random ke-3:", randomizer.Int()) // 8532807521486154107

	fmt.Println("\nUnique Seed\n")
	randomizer2 := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Println("random ke-1:", randomizer2.Int())
	fmt.Println("random ke-2:", randomizer2.Int())
	fmt.Println("random ke-3:", randomizer2.Int())	

	fmt.Println("\nRandom Tipe Data Numerik Lainnya\n")
	fmt.Println("random int:", randomizer2.Int())
	fmt.Println("random float32:", randomizer2.Float32())
	fmt.Println("random uint:", randomizer2.Uint32())	

	fmt.Println("\nAngka Random Index Tertentu (mis:0-99)\n")
	fmt.Println("random int:", randomizer2.Intn(100))

	fmt.Println("\nRandom Tipe Data String\n")
	fmt.Println("random string 5 karakter:", randomString(5))
}

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer3.Intn(len(letters))]
	}
	return string(b)
}
