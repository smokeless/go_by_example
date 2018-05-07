package main

import "fmt"

func ping(pings chan<- string, msg string) {
	//only accepts a channel for sending values.
	//compile time error to recv on this channel.
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	//accept one channel for rcv, one for sends.
	msg := <-pings
	pongs <- msg
}

func channelDirections() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
