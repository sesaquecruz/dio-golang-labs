package main

import (
	"fmt"
	"time"
)

func worker(in chan string, out chan string, text string, delay time.Duration) {
	for msg := range in {
		fmt.Println(msg)
		time.Sleep(delay)
		out <- text
	}
}

func main() {
	in1 := make(chan string)
	in2 := make(chan string)
	out1 := make(chan string)
	out2 := make(chan string)

	go worker(in1, out1, "Ping", time.Duration(1*time.Second))
	go worker(in2, out2, "Pong", time.Duration(1*time.Second))

	var text string

	for i := 0; i < 10; i++ {
		in1 <- "\tack1"
		text = <-out1
		fmt.Println(text)

		in2 <- "\tack2"
		text = <-out2
		fmt.Println(text)
	}

	close(in1)
	close(in2)
	close(out1)
	close(out2)
}
