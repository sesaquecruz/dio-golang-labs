package main

import (
	"fmt"
	game "lab-2/pkg"
	"log"
)

func main() {
	var from, to int

	fmt.Print("from: ")
	_, err := fmt.Scan(&from)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("to: ")
	_, err = fmt.Scan(&to)
	if err != nil {
		log.Fatal(err)
	}

	for i := from; i <= to; i++ {
		text := game.PipPan(i)
		if text != nil {
			fmt.Printf("%d: %s\n", i, *text)
		}
	}
}
