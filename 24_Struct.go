package main

import "fmt"

func inisialisasi(){
	type student struct {
		name string
		grade int
	}

	var s1 = student{}
	s1.name = "wick"
	s1.grade = 2
	
	var s2 = student{"ethan", 2}
	
	var s3 = student{name: "jason"}
	
	fmt.Println("student 1 :", s1.name)
	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 3 :", s3.name)

	var s4 = student{name: "wayne", grade: 2}

	var s5 = student{grade: 2, name: "bruce"}

	fmt.Println("student 4 :", s4.name)
	fmt.Println("student 5 :", s5.name)
}

func pointer(){
	type student struct {
		name string
		grade int
	}
	
	var s1 = student{name: "wick", grade: 2}

	var s2 *student = &s1
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s2.name)
	
	s2.name = "ethan"
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s2.name)
}


func embedded(){
	type person struct {
		name string
		age  int
	}
	
	type student struct {
		grade int
		person
	}
	
	var s1 = student{}
    s1.name = "wick"
    s1.age = 21
    s1.grade = 2

    fmt.Println("name  :", s1.name)
    fmt.Println("age   :", s1.age)
    fmt.Println("age   :", s1.person.age)
    fmt.Println("grade :", s1.grade)
}

func embedded2(){
	type person struct {
		name string
		age  int
	}
	
	type student struct {
		person
		age   int
		grade int
	}

	var s1 = student{}
    s1.name = "wick"
    s1.age = 21        // age of student
    s1.person.age = 22 // age of person

    fmt.Println(s1.name)
    fmt.Println(s1.age)
    fmt.Println(s1.person.age)

	// Pengisian Nilai Sub-Struct
	var p1 = person{name: "wick", age: 21}
	s1 = student{person: p1, grade: 2}
	
	fmt.Println("name  :", s1.name)
	fmt.Println("age   :", s1.age)
	fmt.Println("grade :", s1.grade)
}

func anonym(){
	type person struct {
		name string
		age  int
	}

	var s1 = struct {
        person
        grade int
    }{}
    s1.person = person{"wick", 21}
    s1.grade = 2

    fmt.Println("name  :", s1.person.name)
    fmt.Println("age   :", s1.person.age)
    fmt.Println("grade :", s1.grade)
}

func sliceStruct(){
	type person struct {
		name string
		age  int
	}
	
	var allStudents = []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}
	
	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

}

func anonymSlice(){
	type person struct {
		name string
		age  int
	}

	var allStudents = []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"ethan", 22}, grade: 3},
		{person: person{"bond", 21}, grade: 3},
	}
	
	for _, student := range allStudents {
		fmt.Println(student)
	}
}

func alias(){
	type Person struct {
		name string
		age  int
	}
	type People = Person
	
	var p1 = Person{"wick", 21}
	fmt.Println(p1)
	
	var p2 = People{"wick", 21}
	fmt.Println(p2)


	people := People{"wick", 21}
	fmt.Println(Person(people))
	
	person := Person{"wick", 21}
	fmt.Println(People(person))
	
	
	type People1 struct {
		name string
		age  int
	}
	type People2 = struct {
		name string
		age  int
	}

	// type Number = int
	// var num Number = 12
}

func main() {
    fmt.Println("\nInisialisasi Object Struct\n")
	inisialisasi()
	
	
    fmt.Println("\nVariabel Objek Pointer\n")
	pointer()
	
    fmt.Println("\nEmbedded Struct\n")
	embedded()

	fmt.Println("\nEmbedded Struct Dengan Nama Property Yang Sama\n")
	embedded2()

	fmt.Println("\nAnonymous Struct\n")
	anonym()

	// anonymous struct tanpa pengisian property
		// var s1 = struct {
		// 	person
		// 	grade int
		// }{}

	// anonymous struct dengan pengisian property
		// var s2 = struct {
		// 	person
		// 	grade int
		// }{
		// 	person: person{"wick", 21},
		// 	grade:  2,
		// }
	

	fmt.Println("\nKombinasi Slice & Struct\n")
	sliceStruct()

	fmt.Println("\nInisialisasi Slice Anonymous Struct\n")
	anonymSlice()
	
	// Deklarasi Anonymous Struct Menggunakan Keyword var
	// hanya deklarasi
	// var student struct {
		// 	grade int
		// }
		
		// deklarasi sekaligus inisialisasi
		// var student = struct {
			// 	grade int
			// } {
				// 	12,
				// }
				
				
	// Nested struct
		// type student struct {
		// 	person struct {
		// 		name string
		// 		age  int
		// 	}
		// 	grade   int
		// 	hobbies []string
		// }

		// Deklarasi Dan Inisialisasi Struct Secara Horizontal
		// type person struct { name string; age int; hobbies []string }
		// var p1 = struct { name string; age int } { age: 22, name: "wick" }
		// var p2 = struct { name string; age int } { "ethan", 23 }

	// Tag property dalam struct
	// type person struct {
		// 	name string `tag1`
		// 	age  int    `tag2`
		// }

	fmt.Println("\nType Alias\n")
	alias()
}