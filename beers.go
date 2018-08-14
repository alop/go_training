package main

import (
	"fmt"
)

func main() {
	one := "\xF0\x9F\x8D\xBA"
	two := "\xF0\x9F\x8D\xBB"

	fmt.Println(one, two)
}
