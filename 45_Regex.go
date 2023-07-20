package main

import "fmt"
import "regexp"

func main() {
	// Teknik yang digunakan untuk pencocokan string dengan pola tertentu. Regex biasa dimanfaatkan untuk pencarian dan pengubahan data string
	fmt.Println("\nPenerapan Regexp\n")
    var text = "banana burger soup"
    var regex, err = regexp.Compile(`[a-z]+`)

    if err != nil {
        fmt.Println(err.Error())
    }

    var res1 = regex.FindAllString(text, 2)
    fmt.Printf("%#v \n", res1)
    // []string{"banana", "burger"}

    var res2 = regex.FindAllString(text, -1)
    fmt.Printf("%#v \n", res2)
    // []string{"banana", "burger", "soup"}


	fmt.Println("\nMethod MatchString()\n")
	text = "banana burger soup"
	regex, _ = regexp.Compile(`[a-z]+`)
	
	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)
	// true


	fmt.Println("\nMethod FindString()\n")
	text = "banana burger soup"
	regex, _ = regexp.Compile(`[a-z]+`)
	
	var str = regex.FindString(text)
	fmt.Println(str)
	// "banana"

	fmt.Println("\nMethod FindStringIndex()\n")
	text = "banana burger soup"
	regex, _ = regexp.Compile(`[a-z]+`)
	
	var idx = regex.FindStringIndex(text)
	fmt.Println(idx)
	// [0, 6]
	
	str = text[0:6]
	fmt.Println(str)
	// "banana"	

	fmt.Println("\nMethod FindAllString()\n")
	text = "banana burger soup"
	regex, _ = regexp.Compile(`[a-z]+`)
	
	var str1 = regex.FindAllString(text, -1)
	fmt.Println(str1)
	// ["banana", "burger", "soup"]
	
	var str2 = regex.FindAllString(text, 1)
	fmt.Println(str2)
	// ["banana"]
	
	fmt.Println("\nMethod ReplaceAllString()\n")
	text = "banana burger soup"
	regex, _ = regexp.Compile(`[a-z]+`)
	
	str = regex.ReplaceAllString(text, "potato")
	fmt.Println(str)
	// "potato potato potato"
	
	fmt.Println("\nMethod ReplaceAllStringFunc()\n")
	text = "banana burger soup"
	regex, _ = regexp.Compile(`[a-z]+`)
	
	str = regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	})
	fmt.Println(str)
	// "banana potato soup"
	
	fmt.Println("\nMethod Split()\n")
	text = "banana burger soup"
	regex, _ = regexp.Compile(`[a-b]+`) // split dengan separator adalah karakter "a" dan/atau "b"
	
	var str3 = regex.Split(text, -1)
	fmt.Printf("%#v \n", str3)
	// []string{"", "n", "n", " ", "urger soup"}	
	

}