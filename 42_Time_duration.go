package main

import (
    "fmt"
    "time"
)

func main() {
    start := time.Now()

    time.Sleep(5 * time.Second)

    duration := time.Since(start)

    fmt.Println("time elapsed in seconds:", duration.Seconds())
    fmt.Println("time elapsed in minutes:", duration.Minutes())
    fmt.Println("time elapsed in hours:", duration.Hours())

	fmt.Println("\nKalkulasi Durasi Antara 2 Objek Waktu\n")
	t1 := time.Now()
	time.Sleep(5 * time.Second)
	t2 := time.Now()
	
	duration2 := t2.Sub(t1)
	
	fmt.Println("time elapsed in seconds:", duration2.Seconds())
	fmt.Println("time elapsed in minutes:", duration2.Minutes())
	fmt.Println("time elapsed in hours:", duration2.Hours())	

	// Konversi Angka ke time.Duration
	// var n time.Duration = 5
	// duration := n * time.Second
	// n := 5
	// duration := time.Duration(n) * time.Second		
}