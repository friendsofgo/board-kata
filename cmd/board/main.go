package main

import (
	"fmt"
	"log"

	board "github.com/friendsofgo/board-kata"
)

func main() {
	msg, err := board.ReadInput("data/input.csv")
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Parse each message
	fmt.Println(msg)

	//TODO: Print output into an html file

	fmt.Println("done!")
}
