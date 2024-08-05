package main 
import "fmt"

func main() {
	i:=1
	for i<=3 {
		fmt.Println(i)
		i++
	}

	for i:= range 10 {
		fmt.Println("range",i)
	}

	for n:= range 10{
		if n%2!=0 {
			continue
		}
		fmt.Println(n)
	}
}
