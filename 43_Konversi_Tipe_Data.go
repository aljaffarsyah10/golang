package main

import "fmt"
import "strconv"

func main(){
	fmt.Println("\nKonversi Menggunakan strconv\n")
	fmt.Println("\nFungsi strconv.Atoi()\n")
	var str = "124"
    var num, err = strconv.Atoi(str)

    if err == nil {
        fmt.Println(num) // 124
    }

	fmt.Println("\nFungsi strconv.Itoa()\n")
	var num2 = 124
	var str2 = strconv.Itoa(num2)
	
	fmt.Println(str2) // "124"

	fmt.Println("\nFungsi strconv.ParseInt()\n")
	var str3 = "124"
	var num3, err2 = strconv.ParseInt(str3, 10, 64)
	
	if err2 == nil {
		fmt.Println(num3) // 124
	}
	
	var str4 = "1010"
	var num4, err3 = strconv.ParseInt(str4, 2, 8)
	
	if err3 == nil {
		fmt.Println(num4) // 10
	}
	
	
	fmt.Println("\nFungsi strconv.FormatInt()\n")
	var num5 = int64(24)
	var str5 = strconv.FormatInt(num5, 8)
	
	fmt.Println(str5) // 30	
	
	fmt.Println("\nFungsi strconv.ParseFloat()\n")
	var str6 = "24.12"
	var num6, err4 = strconv.ParseFloat(str6, 32)
	
	if err4 == nil {
		fmt.Println(num6) // 24.1200008392334
	}
	

	fmt.Println("\nFungsi strconv.FormatFloat()\n")
	var num7 = float64(24.12)
	var str7 = strconv.FormatFloat(num7, 'f', 6, 64)
	
	fmt.Println(str7) // 24.120000


	fmt.Println("\nFungsi strconv.ParseBool()\n")
	var str8 = "true"
	var bul, err5 = strconv.ParseBool(str8)
	
	if err5 == nil {
		fmt.Println(bul) // true
	}

	fmt.Println("\nFungsi strconv.FormatBool()\n")
	var bul2 = true
	var str9 = strconv.FormatBool(bul2)
	
	fmt.Println(str9) // true	


	fmt.Println("\nKonversi Data Menggunakan Teknik Casting\n")
	var a float64 = float64(24)
	fmt.Println(a) // 24
	
	var b int32 = int32(24.00)
	fmt.Println(b) // 24	

	fmt.Println("\nCasting string ↔ byte\n")
	var text1 = "halo"
	var b2 = []byte(text1)
	
	fmt.Printf("%d %d %d %d \n", b2[0], b2[1], b2[2], b2[3])
	// 104 97 108 111
	
	var byte1 = []byte{104, 97, 108, 111}
	var s = string(byte1)
	
	fmt.Printf("%s \n", s)
	// halo
	
	var c int64 = int64('h')
	fmt.Println(c) // 104
	
	var d string = string(104)
	fmt.Println(d) // h
	
	
	fmt.Println("\nType Assertions Pada Interface Kosong (interface{})\n")
	// teknik untuk mengambil tipe data konkret dari data yang terbungkus dalam interface{}

	var data = map[string]interface{}{
		"nama":    "john wick",
		"grade":   2,
		"height":  156.5,
		"isMale":  true,
		"hobbies": []string{"eating", "sleeping"},
	}
	
	fmt.Println(data["nama"].(string))
	fmt.Println(data["grade"].(int))
	fmt.Println(data["height"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobbies"].([]string))
	
	// misal tidak tau tipe data interfacenya
	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		case bool:
			fmt.Println(val.(bool))
		case []string:
			fmt.Println(val.([]string))
		default:
			fmt.Println(val.(int))
		}
	}	
}
