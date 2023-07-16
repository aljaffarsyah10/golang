// Eksekusi goroutine bersifat asynchronous, menjadikannya tidak saling tunggu dengan goroutine lain.
// Concurrency atau konkurensi berbeda dengan paralel. Paralel adalah eksekusi banyak proses secara bersamaan. Sedangkan konkurensi adalah komposisi dari sebuah proses. Konkurensi merupakan struktur, sedangkan paralel adalah bagaimana eksekusinya berlangsung.

// Penerapan Goroutine
package main

import "fmt"
import "runtime"

func print(till int, message string) {
    for i := 0; i < till; i++ {
        fmt.Println((i + 1), message)
    }
}

func main() {
    runtime.GOMAXPROCS(2)

    go print(5, "halo")
    print(5, "apa kabar")

    var input string
    fmt.Scanln(&input)

	// Penggunaan Fungsi fmt.Scanln()
	var s1, s2, s3 string
	fmt.Scanln(&s1, &s2, &s3)
	
	// user inputs: "trafalgar d law"
	
	fmt.Println(s1) // trafalgar
	fmt.Println(s2) // d
	fmt.Println(s3) // law	
}