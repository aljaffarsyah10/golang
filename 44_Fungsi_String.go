package main

import "fmt"
import "strings"

func main() {
	fmt.Println("\nFungsi strings.Contains()\n")
    var isExists = strings.Contains("john wick", "wick")
    fmt.Println(isExists)

	fmt.Println("\nFungsi strings.HasPrefix()\n")
	var isPrefix1 = strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1) // true
	
	var isPrefix2 = strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2) // false
	
	fmt.Println("\nFungsi strings.HasSuffix()\n")
	var isSuffix1 = strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix1) // false
	
	var isSuffix2 = strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix2) // true
	
	fmt.Println("\nFungsi strings.Count()\n")
	var howMany = strings.Count("ethan hunt", "t")
	fmt.Println(howMany) // 2
	
	fmt.Println("\nFungsi strings.Index()\n")
	var index1 = strings.Index("ethan hunt", "ha")
	fmt.Println(index1) // 2
	// Jika diketemukan dua substring, maka yang diambil adalah yang pertama	
	var index2 = strings.Index("ethan hunt", "n")
	fmt.Println(index2) // 4
	
	fmt.Println("\nFungsi strings.Replace()\n")
	var text = "banana"
	var find = "a"
	var replaceWith = "o"
	
	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) // "bonana"
	
	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // "bonona"
	
	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3) // "bonono"
	
	
	fmt.Println("\nFungsi strings.Repeat()\n")
	var str = strings.Repeat("na", 4)
	fmt.Println(str) // "nananana"
	
	fmt.Println("\nFungsi strings.Split()\n")
	var string1 = strings.Split("the dark knight", " ")
	fmt.Println(string1) // output: ["the", "dark", "knight"]
	
	var string2 = strings.Split("batman", "")
	fmt.Println(string2) // output: ["b", "a", "t", "m", "a", "n"]
	
	fmt.Println("\nFungsi strings.Join()\n")
	var data = []string{"banana", "papaya", "tomato"}
	var str2 = strings.Join(data, "-")
	fmt.Println(str2) // "banana-papaya-tomato"

	fmt.Println("\nFungsi strings.ToLower()\n")
	var str3 = strings.ToLower("aLAy")
	fmt.Println(str3) // "alay"
	
	fmt.Println("\nFungsi strings.ToUpper()\n")
	var str4 = strings.ToUpper("eat!")
	fmt.Println(str4) // "EAT!"	
}