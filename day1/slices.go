package main
import (
	"fmt"
	"slices"
)

func main () {
	a := [...]int{1,2,3,4,5}
	fmt.Println(a[2:4])
	fmt.Println(a[1])

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	s = append(s, "d")
	s = append(s, "e", "f")
	b := []string{"a","b","c"}
	c := []string{"a","b","c"}
	fmt.Println(slices.Equal(b,c))
}
