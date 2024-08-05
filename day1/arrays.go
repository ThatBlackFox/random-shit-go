package main
import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)
	a[4] = 4000000
	fmt.Println("set:",a)
	fmt.Println("lao: the chinese hacker: ",a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1,2,3,4,5}
	fmt.Println(b)
	b = [...]int{420,69,109,40,0}
	twoD := [2][3]int {
		{1,2,3},
		{4,5,6},
	}
	fmt.Println(twoD)
}
