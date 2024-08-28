package main

import (
	"fmt"
)

func sayMoo(until int) {
	if until > 0 {
		fmt.Println("Moo")
		until--
		sayMoo(until)
	}
}

func main() {
	sayMoo(10)
}
