package main

import "fmt"

func main() {
	nums := []int{2,44,9}
	sum := 0
	for _, num := range nums {
		sum+=num
	}
	fmt.Println(sum)

	sum = 0
	for num := range len(nums) {
		sum += num
	}
	fmt.Println(sum)

	dict := make(map[string]int)
	dict["Chipi Chipi Yale"] = 200*10+4
	dict["Diablo Of Babilon"] = 25-1
	
	for k,v := range dict {
		fmt.Printf("%s -> %d\n",k,v)
	}

}
