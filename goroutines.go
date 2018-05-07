package main

import "fmt"

func myFunction(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func goRoutine() {
	myFunction("directCall") //synchronously

	go myFunction("goroutine") //concurrent execution with calling function.

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")
}
