package main

// import "belajar-golang-level-akses/library"
// import "fmt"

// Import Dengan Prefix Tanda Titik
import ."belajar-golang-level-akses/library"
// Pemanfaatan Alias Ketika Import Package
import f "fmt"


func main(){
    // library.SayHello("ethan")
    SayHello("ethan")

	// var s1 = library.Student{"ethan", 21}
	var s1 = Student{"ethan", 21}
    f.Println("name", s1.Name)
    f.Println("grade", s1.Grade)

	sayHello2("ethan")

	f.Printf("Name  : %s\n", Student2.Name)
    f.Printf("Grade : %d\n", Student2.Grade)
}