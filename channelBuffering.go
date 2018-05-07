package main

import "fmt"

func goChannelBuffering() {
	messages := make(chan string, 2) //new channel takes 2 values

	messages <- "buffered" //sends channel <- "thing"
	messages <- "channel"
	//because this is buffered we can send values without a concurrent
	//recv.
	fmt.Println(<-messages) //rcv <-channel
	fmt.Println(<-messages)
}
