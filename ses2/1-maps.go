package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)
	fmt.Println(m)

	m["A key to memes"] = 0
	fmt.Println(m)

	n := make(map[int]string)
	fmt.Println(n)
	n[0] = "basically a bad array"

	if maps.Equal(n,n) {
		fmt.Println("hi")
	} else {
		fmt.Println("hehe")
	}

}
