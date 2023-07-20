package main

import "fmt"
import "os"
// import "flag"

func main() {
//	Run satu satu sesuai argumen ataupun flag

	var argsRaw = os.Args
    fmt.Printf("-> %#v\n", argsRaw)
    // -> []string{".../bab45", "banana", "potato", "ice cream"}

    var args = argsRaw[1:]
    fmt.Printf("-> %#v\n", args)
    // -> []string{"banana", "potatao", "ice cream"}

	// Untuk eksekusinya sendiri bisa menggunakan go run ataupun dengan cara build-execute.
	// Menggunakan go run
	//  go run bab45.go banana potato "ice cream"
	// Menggunakan go build
	//  go build bab45.go
	//  $ ./bab45 banana potato "ice cream"	


	// fmt.Println("\nPenggunaan Flag\n")
    // var name = flag.String("name", "anonymous", "type your name")
    // var age = flag.Int64("age", 25, "type your age")

    // flag.Parse()
    // fmt.Printf("name\t: %s\n", *name)
    // fmt.Printf("age\t: %d\n", *age)

	// Cara penulisan arguments menggunakan flag:
	// go run bab45.go -name="john wick" -age=28	

	// fmt.Println("\nDeklarasi Flag Dengan Cara Passing Reference Variabel Penampung Data\n")
	// // cara ke-1
	// var data1 = flag.String("name", "anonymous", "type your name")
	// fmt.Println(*data1)

	// // cara ke-2
	// var data2 string
	// flag.StringVar(&data2, "gender", "male", "type your gender")
	// fmt.Println(data2)	
	
}