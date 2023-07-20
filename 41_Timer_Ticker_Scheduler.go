package main

import "fmt"
import "os"
import "time"

func main () {
	fmt.Println("\nFungsi time.Sleep()\n")
    fmt.Println("start")
    time.Sleep(time.Second * 4)
    fmt.Println("after 4 seconds")

    fmt.Println("\nScheduler Menggunakan time.Sleep()\n")
	for i := 1; i <= 3; i++ {
        fmt.Println("Hello !!")
        time.Sleep(1 * time.Second)
    }

    fmt.Println("\nFungsi time.NewTimer()\n")
    var timer = time.NewTimer(4 * time.Second)
    fmt.Println("start")
    <-timer.C
    fmt.Println("finish")
    
    fmt.Println("\nFungsi time.AfterFunc()\n")
    var ch = make(chan bool)

    time.AfterFunc(4*time.Second, func() {
        fmt.Println("expired")
        ch <- true
    })
    
    fmt.Println("start")
    <-ch
    fmt.Println("finish")
    
    // Fungsi time.After()
    <-time.After(4 * time.Second)
    fmt.Println("expired")

    fmt.Println("\nKombinasi Timer & Goroutine\n")
    var timeout = 5
    var ch2 = make(chan bool)

    go timer2(timeout, ch2)
    go watcher(timeout, ch2)

    var input string
    fmt.Print("what is 725/25 ? ")
    fmt.Scan(&input)

    if input == "29" {
        fmt.Println("the answer is right!")
    } else {
        fmt.Println("the answer is wrong!")
    }

    fmt.Println("\nScheduler Menggunakan Ticker\n")
    done := make(chan bool)
    ticker := time.NewTicker(time.Second)

    go func() {
        time.Sleep(3 * time.Second) // wait for 10 seconds
        done <- true
    }()

    for {
        select {
        case <-done:
            ticker.Stop()
            return
        case t := <-ticker.C:
            fmt.Println("Hello !!", t)
        }
    }
}

func timer2(timeout int, ch chan<- bool) {
    time.AfterFunc(time.Duration(timeout)*time.Second, func() {
        ch <- true
    })
}

func watcher(timeout int, ch <-chan bool) {
    <-ch
    fmt.Println("\ntime out! no answer more than", timeout, "seconds")
    os.Exit(0)
}