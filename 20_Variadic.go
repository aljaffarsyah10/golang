package main

import "fmt"
import "strings"

func main() {
    var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
    var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
    fmt.Println(msg)

	// Pengisian Parameter Fungsi Variadic Menggunakan Data Slice
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	avg = calculate(numbers...)
	msg = fmt.Sprintf("Rata-rata : %.2f", avg)

	// Fungsi Dengan Parameter Biasa & Variadic
	yourHobbies("wick", "sleeping", "eating")

fmt.Println(msg)
}

func calculate(numbers ...int) float64 {
    var total int = 0
    for _, number := range numbers {
        total += number
    }

    var avg = float64(total) / float64(len(numbers))
    return avg
}

func yourHobbies(name string, hobbies ...string) {
    var hobbiesAsString = strings.Join(hobbies, ", ")

    fmt.Printf("Hello, my name is: %s\n", name)
    fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}