package main

import "fmt"

type employee struct {
	firstName string
	lastName  string
	age       int
	weight    int
}

func main() {
	fred := employee{firstName: "Fred",
		age: 45, weight: 156}

	fmt.Println(fred)
}
