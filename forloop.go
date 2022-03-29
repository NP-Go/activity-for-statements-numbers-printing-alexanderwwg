package main

import "fmt"

var number1 int

func main() {
	var input int
	number1 = 1000
	fmt.Println("Input a command between 1 to 3.")
	fmt.Scanln(&input)
	runTest(input)
}
func runTest(input int) {
	switch input {
	case 1:
		forLoopCount(number1)
	case 2:
		forLoopEvenCount(number1)
	case 3:
		forLoopOddCount(number1)
	default:
		fmt.Println("Error, wrong number input.")
		break
	}
}
func forLoopCount(number int) {
	fmt.Println("Printing all numbers.")
	for i := 0; i <= number; i++ {
		fmt.Println(i)
	}
}
func forLoopEvenCount(number int) {
	fmt.Println("Printing all even numbers.")
	// do not include 0
	for i := 1; i <= number; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
func forLoopOddCount(number int) {
	fmt.Println("Printing all odd numbers.")
	for i := 0; i <= number; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}
