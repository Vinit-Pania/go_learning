package main

import (
	"fmt"
)

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

	//Learning Slice which is wrapper of aare

	int_Slice1 := []int32{2, 21, 221, 2221, 2221}
	fmt.Println(int_Slice1)
	fmt.Printf("The lebgth of the aare is %v and the capacity of the array is %v", len(int_Slice1), cap(int_Slice1))

	int_Slice1 = append(int_Slice1, 222211)
	fmt.Println(int_Slice1)
	fmt.Printf("\n The lebgth of the aare is %v and the capacity of the array is %v", len(int_Slice1), cap(int_Slice1))

	//cant reAD THE value from the added part of the capacity
	fmt.Println(int_Slice1[4])

	var int_Slice2 []int32 = []int32{222222211, 2121212212}
	int_Slice2 = append(int_Slice1, int_Slice2...)
	fmt.Println(int_Slice2)

	//another way makeing slice using slice which has lemgth and capacity default length == capacity
	var int_Slice3 []int32 = make([]int32, 1, 10)
	fmt.Println(int_Slice3)

	//learinig map which is key:value pair
	var myMap1 map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap1)

	var myMap2 map[string]uint8 = map[string]uint8{"Adam": 12, "Eve": 25}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Json"]) // will give default value which is 0

	//has another valve which is a boolean if value present then true else false
	var age, ok = myMap2["Adam"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("Invalid namu")
	}

	age, ok = myMap2["Json"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("Invalid namu")
	}

	//delete value from map use buitin fn called delete
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	delete(myMap2, "Eve")
	age, ok = myMap2["Eve"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("Invalid namu")
	}

	// for loops in map
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	var Hexenzirkel map[string]uint8 = map[string]uint8{"alice": 1, "Barbeloth": 2, "I. Ivanovna N.": 3, "Andersdotter": 4, "Nicole Reeyn": 5, "Octavia ": 6, "Rhinedottir": 7}
	for namae := range Hexenzirkel {
		fmt.Printf("Name: %v\n", namae)
	}

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	for namae, position := range Hexenzirkel {
		fmt.Printf("Name: %v , Position : %v \n ", namae, position)
	}

	//For loop in Array or slice
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	for i, v := range intArr {
		fmt.Printf("Index : %v , Range : %v \n", i, v)
	}

	//While loop not directly present we can makeshift using for
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
		for {
			if i >= 10 {
				fmt.Println("break")
				break
			}
			fmt.Println(i)
			i = i + 1
		}
	}

	//Differernt method of while loop
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
