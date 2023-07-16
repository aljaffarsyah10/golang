package main

import (
	"errors"
    "fmt"
    "strconv"
	"strings"
)

func main() {
	fmt.Println("\nPemanfaatan Error\n")
    var input string
    fmt.Print("Type some number: ")
    fmt.Scanln(&input)

    var number int
    var err error
    number, err = strconv.Atoi(input)

    if err == nil {
        fmt.Println(number, "is number")
    } else {
        fmt.Println(input, "is not number")
        fmt.Println(err.Error())
    }


	fmt.Println("\nMembuat Custom Error\n")
	var name string
    fmt.Print("Type your name: ")
    fmt.Scanln(&name)

    if valid, err := validate(name); valid {
        fmt.Println("halo", name)
    } else {
        fmt.Println(err.Error())
    }


	fmt.Println("\nPenggunaan panic\n")
	// menampilkan stack trace error sekaligus menghentikan flow goroutine (karena main() juga merupakan goroutine, maka behaviour yang sama juga berlaku)
	// proses akan terhenti, apapun setelah tidak di-eksekusi kecuali proses yang sudah di-defer sebelumnya (akan muncul sebelum panic error).
    // var name string
    fmt.Print("Type your name: ")
    fmt.Scanln(&name)

    if valid, err := validate(name); valid {
        fmt.Println("halo", name)
    } else {
        panic(err.Error())
        fmt.Println("end")
    }


	fmt.Println("\nPenggunaan recover\n")
	defer catch()

    // var name string
    fmt.Print("Type your name: ")
    fmt.Scanln(&name)

    if valid, err := validate(name); valid {
        fmt.Println("halo", name)
    } else {
        panic(err.Error())
        fmt.Println("end")
    }


	fmt.Println("\nPemanfaatan recover pada IIFE\n")
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Panic occured", r)
        } else {
            fmt.Println("Application running perfectly")
        }
    }()

    panic("some error happen")


	data := []string{"superman", "aquaman", "wonder woman"}

    for _, each := range data {

        func() {

            // recover untuk IIFE dalam perulangan
            defer func() {
                if r := recover(); r != nil {
                    fmt.Println("Panic occured on looping", each, "| message:", r)
                } else {
                    fmt.Println("Application running perfectly")
                }
            }()

            panic("some error happen")
        }()

    }

}

func validate(input string) (bool, error) {
    if strings.TrimSpace(input) == "" {
        return false, errors.New("cannot be empty")
    }
    return true, nil
}

func catch() {
    if r := recover(); r != nil {
        fmt.Println("Error occured", r)
    } else {
        fmt.Println("Application running perfectly")
    }
}