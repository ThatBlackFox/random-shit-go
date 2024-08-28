package main

import (
	"fmt"
)

func IntCounter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	counter := IntCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
