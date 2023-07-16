package main

import "fmt"

func main(){

	// Penggunaan Map
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei",     chicken["mei"])     // mei 0

	// Inisialisasi Nilai Map
	// var data map[string]int 
	// data["one"] = 1
	// akan muncul error!

	var data = map[string]int{}
	data["one"] = 1
	// tidak ada error

	// cara horizontal
	// var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// cara vertical
	// var chicken2 = map[string]int{
	//     "januari":  50,
	//     "februari": 40,
	// }

	// Iterasi Item Map Menggunakan for - range
	chicken = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}

	// Menghapus Item Map
	chicken = map[string]int{"januari": 50, "februari": 40}

	fmt.Println(len(chicken)) // 2
	fmt.Println(chicken)

	delete(chicken, "januari")

	fmt.Println(len(chicken)) // 1
	fmt.Println(chicken)

	// Deteksi Keberadaan Item Dengan Key Tertentu
	chicken = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	// Kombinasi Slice & Map
	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue",   "gender": "male"},
		map[string]string{"name": "chicken red",    "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	chickens = []map[string]string{
		{"name": "chicken blue",   "gender": "male"},
		{"name": "chicken red",    "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	// var data2 = []map[string]string{
	// 	{"name": "chicken blue", "gender": "male", "color": "brown"},
	// 	{"address": "mangga street", "id": "k001"},
	// 	{"community": "chicken lovers"},
	// }

}
