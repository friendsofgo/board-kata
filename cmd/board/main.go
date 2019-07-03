package main

import (
	"fmt"
	"log"

	board "github.com/friendsofgo/board-kata"
	"github.com/friendsofgo/board-kata/parser"
)

func main() {
	msg, err := board.ReadInput("data/input.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Parse each message
	for _, line := range msg {
		translated := parser.Parse(line)
		fmt.Println(translated)
	}

	//TODO: Print output into an html file

	fmt.Println("done!")
}
