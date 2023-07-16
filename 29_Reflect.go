// Mencari Tipe Data & Value Menggunakan Reflect
package main

import "fmt"
import "reflect"

func main() {
    var number = 23
    var reflectValue = reflect.ValueOf(number)

    fmt.Println("tipe  variabel :", reflectValue.Type())

    if reflectValue.Kind() == reflect.Int {
        fmt.Println("nilai variabel :", reflectValue.Int())
    }

	// Pengaksesan Nilai Dalam Bentuk interface{}
	fmt.Println("tipe  variabel :", reflectValue.Type())
	fmt.Println("nilai variabel :", reflectValue.Interface())
	
	var nilai = reflectValue.Interface().(int)
	
	fmt.Println("nilai variabel :", nilai)


	// Pengaksesan Informasi Property Variabel Objek
	var s1 = &student{Name: "wick", Grade: 2}
    s1.getPropertyInfo()

	// Pengaksesan Informasi Method Variabel Objek
	s1 = &student{Name: "john wick", Grade: 2}
    fmt.Println("nama :", s1.Name)

    reflectValue = reflect.ValueOf(s1)
    var method = reflectValue.MethodByName("SetName")
    method.Call([]reflect.Value{
        reflect.ValueOf("wick"),
    })

    fmt.Println("nama :", s1.Name)
}

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama      :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai     :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func (s *student) SetName(name string) {
    s.Name = name
}