package main

import (
	"fmt"
)

func main() {
	maths := map[string][]int{
		"a": {4, 5},
		"d": {0, 3},
		"m": {4, 8},
		"s": {1, 4},
	}

	fmt.Println("Matrix at start\n", maths)

	for f, xy := range maths {
		maths[f] = append(maths[f], doMath(f, xy))
	}
	fmt.Println("Matrix at end\n", maths)
}

func doMath(op string, nums []int) (ans int) {
	switch op {
	case "a":
		ans = add(nums[0], nums[1])
	case "s":
		ans = sub(nums[0], nums[1])
	case "m":
		ans = mul(nums[0], nums[1])
	case "d":
		ans = div(nums[0], nums[1])
	}
	return
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func sub(num1 int, num2 int) int {
	return num1 - num2
}

func div(num1 int, num2 int) int {
	return num1 / num2
}

func mul(num1 int, num2 int) int {
	return num1 * num2
}
