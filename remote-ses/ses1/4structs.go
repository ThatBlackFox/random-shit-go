package main

import "fmt"

type meme struct {
	name   string
	rating float32
	isdark bool
}

func main() {

	nigeshMeme := meme{name: "Don't try to play fool with me nigesh"}
	nigeshMeme.isdark = true
	nigeshMeme.rating = 9.782

	fmt.Println("My nigesh")
	fmt.Println(nigeshMeme)
}
