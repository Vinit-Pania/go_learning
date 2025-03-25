package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intnum int = 12
	fmt.Println(intnum)

	var uintNum uint8 = 123
	fmt.Println(uintNum) //Unsigned num ie positive numbers

	var floatNum32 float32 = 1112223.121
	fmt.Println(floatNum32)

	var floatnum float64 = 121211121
	fmt.Println(floatnum)

	var floatall = floatnum + float64(intnum)
	fmt.Println(floatall)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2) // dividend
	fmt.Println(intNum1 % intNum2) // remainder

	var booleanum bool = true
	fmt.Println(booleanum) // default value is false

	var myString string = "Eblo World"
	fmt.Println(myString)

	var myStingHasSPace string = `My String Has
	Space`
	fmt.Println(myStingHasSPace)

	var myStringCanConatinate string = "Eblo" + " " + "World"
	fmt.Println(myStringCanConatinate)

	fmt.Println(len("test")) // Provide in bytes of that string not the string

	fmt.Println(utf8.RuneCountInString("test")) // Actually provide the  length of the test

	var myRune rune = 'a'
	fmt.Println(myRune) //Provide unicode of letter a

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst string = "A const value "
	fmt.Println(myConst)

	const pi float32 = 3.141
	fmt.Println(pi)

}
