package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error, specify start and end")
		os.Exit(1)
	}

	var start, _ = strconv.Atoi(os.Args[1])
	var end, _ = strconv.Atoi(os.Args[2])

	fmt.Println(getSlice(start, end))
}

func getSlice(_start int, _end int) []int {

	var seq = []int{0, 1}
	var requested []int
	current := seq[len(seq)-1]
	previous := seq[len(seq)-2]
	next := current + previous
	seq = append(seq, next)
	for next < 100000 {
		current := seq[len(seq)-1]
		previous := seq[len(seq)-2]
		next = current + previous
		seq = append(seq, next)
	}
	for _, value := range seq {
		if value >= _start && value <= _end {
			requested = append(requested, value)
		}
	}
	return requested
}
