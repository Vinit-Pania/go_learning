package main

import (
	"errors"
	"fmt"
)

func main() {
	printSomething()

	var printValue string = "test is the new way"
	printWithStringParameter(printValue)

	var numerator int = 12
	var denominator int = 13

	var result, remainder, err = intDivision(numerator, denominator)

	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v and remainder is %v", result, remainder)
	}

	switch {
	case err != err:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result is %v", result)
	default:
		fmt.Printf("Ther result is %v and the remainder is %v", result, remainder)
	}

	// conditional switch statment

	switch remainder {
	case 0:
		fmt.Printf("The division wAS exact")
	case 1, 2:
		fmt.Printf("The division was close")
	default:
		fmt.Println("Its not even close")
	}

}

func printSomething() {
	fmt.Println("hi boiz")
}

func printWithStringParameter(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("value cant be zero")
	}
	var result = numerator / denominator
	var remainder = numerator % denominator
	return result, remainder, err
}
