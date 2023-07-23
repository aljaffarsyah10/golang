// Vendoring di Go merupakan kapabilitas untuk mengunduh semua dependency atau 3rd party, untuk disimpan di lokal dalam folder project, dalam folder bernama vendor.

// di terminal : 
// mkdir belajar-vendor
// cd belajar-vendor
// go mod init belajar-vendor
// go get -u github.com/novalagung/gubrak/v2

// jalankan go mod vendor

package main

import (
    "fmt"
    gubrak "github.com/novalagung/gubrak/v2"
)

func main() {
    fmt.Println(gubrak.RandomInt(10, 20))
}