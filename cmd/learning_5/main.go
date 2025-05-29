// String Rune and Byte

package main

import (
	"fmt"
	"strings"
)

func main() {
	//Strings does not index the actual char but the byte of the char to slove this use rune
	var myString = []rune("resume")
	var indexed = myString[1]

	fmt.Println(myString)
	fmt.Println(indexed)
	fmt.Printf("%v,%T", indexed, indexed)

	for g, v := range myString {
		fmt.Println(g, v)
	}

	var myRune = 'a'
	fmt.Printf("\nMy Rune = %v", myRune)

	//String building using concact

	var strSlice = []string{"a", "e", "i", "o", "u"}
	var catStr = ""

	//println("slice string =", strSlice)
	for i := range strSlice {
		catStr += strSlice[i]
		fmt.Printf("\n concat str = %v", catStr)
	}

	//After concacat we cannot change so its immmutable
	//for manuplating string we use built in package called "strings" which has string builder
	var strBuilder strings.Builder

	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
		var cattedStr = strBuilder.String()
		fmt.Printf("\n%v", cattedStr)
	}
}
