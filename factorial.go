package main

import (
	"fmt"
	"os"
)

func main() {
	var num = getInput()
	calculateFactorial(num)

}
func getInput() int {
	var in int
	fmt.Println("Enter a number between 1 - 100")
	fmt.Scanln(&in)

	return in
}

func askAgain() {
	e := recover()
	fmt.Println("recovered from", e)
	var retry_num int
	retry_num = getInput()

	calculateFactorial(retry_num)
}

func calculateFactorial(_num int) {
	defer askAgain()

	if _num < 0 || _num > 100 {
		panic("Number out of range")
	} else {
		var product float64
		product = 1
		var factors []int

		for factor := 1; factor <= _num; factor++ {
			factors = append(factors, factor)
		}

		for _, value := range factors {
			product = product * float64(value)
		}

		fmt.Println(_num, "! is ", product)
		os.Exit(0)
	}
}
