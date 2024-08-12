package main

import (
	"fmt"
)

func squish(nums ...int) int {
	total := 0
	for _, i := range nums {
		total+=i
	}
	return total
}

func main() {
	to_spread := []int{1,2,3,4}
	fmt.Println("sum:",squish(to_spread...))
}
