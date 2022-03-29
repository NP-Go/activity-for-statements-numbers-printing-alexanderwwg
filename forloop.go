package main

import "fmt"

var number1 int
var rangeinput1, rangeinput2 int

func main() {
	var input int
	fmt.Println("Input number 1:")
	fmt.Scanln(&rangeinput1)
	fmt.Println("Input number 2:")
	fmt.Scanln(&rangeinput2)

	fmt.Println("Input a command between 1 to 3. \n 1 = print all numbers.\n2= print all even numbers.\n3= print all odd numbers.")
	fmt.Scanln(&input)
	runTest(input)
}

// run the inputs from the scanln
func runTest(input int) {
	min, max := getMinAndMaxValue(rangeinput1, rangeinput2)

	fmt.Printf("Min : %v", min)
	fmt.Printf("Max : %v", max)
	switch input {
	case 1:
		forLoopCount(min, max)
	case 2:
		forLoopEvenCount(min, max)
	case 3:
		forLoopOddCount(min, max)
	default:
		fmt.Println("Error, wrong number input.")
	}
}

// case 1 to 3
func forLoopCount(lowerRange, upperRange int) {
	fmt.Println("Printing all numbers.")
	for i := lowerRange; i <= upperRange; i++ {
		fmt.Println(i)
	}
}
func forLoopEvenCount(lowerRange, upperRange int) {
	fmt.Println("Printing all even numbers.")
	// do not include 0
	for i := lowerRange; i <= upperRange; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
func forLoopOddCount(lowerRange, upperRange int) {
	fmt.Println("Printing all odd numbers.")
	for i := lowerRange; i <= upperRange; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}

// gets max min ints
func getMinAndMaxValue(number1, number2 int) (min, max int) {
	if number1 > number2 {
		number1, number2 = number2, number1
	}
	min = number1
	max = number2
	return min, max
}
