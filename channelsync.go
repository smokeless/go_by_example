package main

import "fmt"
import "time"

func worker(done chan bool) { //takes done of type channel of type bool
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func channelSync() {
	done := make(chan bool, 1) //new channel takes 1 bool
	go worker(done)
	<-done //done channel is used to notify another channel we're done
	//why is done not a pass by reference?
}
