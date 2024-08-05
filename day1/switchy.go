package main
import (
	"fmt"
	"time"
)

func main() {
	i:=2
	fmt.Println("Aise ", i, "ko likho ")
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("Gaand me danda le~")
	}

	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("Off duty par call nahi karne ka bnchod")
	default:
		fmt.Println("Yes, Manager?")
	}
}

