package main

import "fmt"

func main() {
	t1 := total(1, 5, 9, 3)
	fmt.Println(t1)
}

func total(values ...int) int {
	var result int
	for _, item := range values {
		result += item
	}

	return result
}
