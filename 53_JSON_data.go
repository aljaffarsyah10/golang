package main

import "encoding/json"
import "fmt"

// Decode JSON Ke Variabel Objek Struct
type User struct {
    FullName string `json:"Name"`
    Age      int
}

// Decode JSON Ke map[string]interface{} & interface{}

func main() {
	fmt.Println("\nDecode JSON Ke Variabel Objek Struct\n")
    var jsonString = `{"Name": "john wick", "Age": 27}`
    var jsonData = []byte(jsonString)

    var data User

    var err = json.Unmarshal(jsonData, &data)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("user :", data.FullName)
    fmt.Println("age  :", data.Age)


	fmt.Println("\nDecode JSON Ke map[string]interface{} & interface{}\n")
	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)
	
	fmt.Println("user :", data1["Name"])
	fmt.Println("age  :", data1["Age"])	

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)
	
	var decodedData = data2.(map[string]interface{})
	fmt.Println("user :", decodedData["Name"])
	fmt.Println("age  :", decodedData["Age"])	


	fmt.Println("\nDecode Array JSON Ke Array Objek\n")
	var jsonString2 = `[
		{"Name": "john wick", "Age": 27},
		{"Name": "ethan hunt", "Age": 32}
	]`
	
	var data3 []User
	
	var err2 = json.Unmarshal([]byte(jsonString2), &data3)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}
	
	fmt.Println("user 1:", data3[0].FullName)
	fmt.Println("user 2:", data3[1].FullName)	


	fmt.Println("\nEncode Objek Ke JSON String\n")
	var object = []User{{"john wick", 27}, {"ethan hunt", 32}}
	var jsonData2, err3 = json.Marshal(object)
	if err3 != nil {
		fmt.Println(err3.Error())
		return
	}
	
	var jsonString3 = string(jsonData2)
	fmt.Println(jsonString3)	
}