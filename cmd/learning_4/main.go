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

	//Learning Slice which is wrapper of aare

	int_Slice1 := []int32{2, 21, 221, 2221, 2221}
	fmt.Println(int_Slice1)
	fmt.Println("The lebgth of the aare is %v and the capacity of the array is %v", len(int_Slice1), cap(int_Slice1))

	int_Slice1 = append(int_Slice1, 222211)
	fmt.Println(int_Slice1)
	fmt.Println("\n The lebgth of the aare is %v and the capacity of the array is %v", len(int_Slice1), cap(int_Slice1))

	//cant reAD THE value from the added part of the capacity
	fmt.Println(int_Slice1[4])

	var int_Slice2 []int32 = []int32{222222211, 2121212212}
	int_Slice2 = append(int_Slice1, int_Slice2...)
	fmt.Println(int_Slice1)

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
		fmt.Println("The age is %v", age)
	} else {
		fmt.Println("Invalid namu")
	}

	age, ok = myMap2["Json"]
	if ok {
		fmt.Println("The age is %v", age)
	} else {
		fmt.Println("Invalid namu")
	}

	//delete value from map use buitin fn called delete
	delete(myMap2, "Eve")
	age, ok = myMap2["Eve"]
	if ok {
		fmt.Println("The age is %v", age)
	} else {
		fmt.Println("Invalid namu")
	}
}
