package main

import (
	"fmt"
)

// TODO: Implement relaying of message with Channel Direction

func genMsg(ch01 chan<- string) {
	// send message on ch1
	ch01 <- "mensagem"
}

func relayMsg(ch01 <-chan string, ch02 chan<- string) {
	// recv message on ch1
	m:= <-ch01
	// send it on ch2
	ch02 <- m
}

func main() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)

	// spine goroutine genMsg and relayMsg
	go genMsg(ch1)

	go relayMsg(ch1, ch2)

	// recv message on ch2

	v := <-ch2

	fmt.Println(v)
}
