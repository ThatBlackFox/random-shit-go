package main

import (
	"fmt"
)

func vals(a,b int)(int, int) {
	return a,b
}

func main() {
	a, b := vals(1337,911)
	fmt.Println(a,b)
}
