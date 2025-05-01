package main

import "fmt"

func main() {

	//array has FIXED LENGTH, SAME TYPE, INDEXABLE and Continous in Memeory

	var intArr [3]int32 //defalut initialiZation is 32bits
	intArr[1] = 1212
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0]) //checking memory location of an array
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intArr3 [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr3[2])
	fmt.Println(intArr3)

	intArr4 := [4]int32{1, 12, 112, 1112}
	fmt.Println(intArr4)

	intArr5 := [...]int32{1, 1122, 111222, 1112222, 111122222}
	fmt.Println(intArr5)

}
