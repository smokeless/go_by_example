package main

import "fmt"

/*
Channels are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine and
receive those values into another goroutine.
*/
func goChannels() {
	messages := make(chan string) //chan valtype

	go func() { messages <- "ping" }() //value into channel

	msg := <-messages
	fmt.Println(msg)

	/*
		By default sends and receives block until both the sender
		 and receiver are ready. This property allowed us to
		 wait at the end of our program for the "ping" message
		 without having to use any other synchronization.
	*/
}
