package library

import "fmt"

// Penggunaan Package, Import, Dan Hak Akses Exported dan Unexported
func SayHello(name string){
	fmt.Println("Hello")
	introduce(name)
}

func introduce(name string){
	fmt.Println("nama saya",name)
}

// Penggunaan Hak Akses Exported dan Unexported pada Struct dan Propertinya
type Student struct {
    Name  string
    Grade int
}

// Fungsi init()
var Student2 = struct {
    Name  string
    Grade int
}{}

func init() {
    Student2.Name = "John Wick"
    Student2.Grade = 2

    fmt.Println("--> library/library.go imported")
}