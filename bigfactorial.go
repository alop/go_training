package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	var num = getInput()
	calculateFactorial(num)

}
func getInput() int64 {
	var in int64
	fmt.Println("Enter a number between 1 - 100")
	fmt.Scanln(&in)

	return in
}

func askAgain() {
	e := recover()
	fmt.Println("recovered from", e)
	var retry_num int64
	retry_num = getInput()

	calculateFactorial(retry_num)
}

func calculateFactorial(_num int64) {
	defer askAgain()

	if _num < 0 || _num > 100 {
		panic("Number out of range")
	} else {
		product := big.NewInt(1)
		product.MulRange(1, _num)
		fmt.Println(_num, "! is ", product)
		os.Exit(0)
	}
}
