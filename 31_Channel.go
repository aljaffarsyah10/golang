package main

import "fmt"
import "runtime"

func printMessage(what chan string) {
    fmt.Println(<-what)
}

func main() {

    fmt.Println("\nPenerapan Channel\n")
    runtime.GOMAXPROCS(2)

    var messages = make(chan string)

    var sayHelloTo = func(who string) {
        var data = fmt.Sprintf("hello %s", who)
        messages <- data
    }

    go sayHelloTo("john wick")
    go sayHelloTo("ethan hunt")
    go sayHelloTo("jason bourne")

    var message1 = <-messages
    fmt.Println(message1)

    var message2 = <-messages
    fmt.Println(message2)

    var message3 = <-messages
    fmt.Println(message3)

    fmt.Println("\nChannel Sebagai Tipe Data Parameter\n")
	// var messages = make(chan string)

    for _, each := range []string{"wick", "hunt", "bourne"} {
        go func(who string) {
            var data = fmt.Sprintf("hello %s", who)
            messages <- data
        }(each)
    }

    for i := 0; i < 3; i++ {
        printMessage(messages)
    }

	// Iterasi Data Slice/Array Langsung Pada Saat Inisialisasi
	// for _, each := range []string{"wick", "hunt", "bourne"} {
	// 	// ...
	// }

    fmt.Println("\nEksekusi Goroutine Pada IIFE\n")
	// var messages = make(chan string)

	go func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}("wick")
	
	var message = <-messages
	fmt.Println(message)
}