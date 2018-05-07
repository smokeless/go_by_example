package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

/*
This function intSeq returns another function,
which we define anonymously in the body of intSeq.
The returned function closes over the variable i to form a closure.
*/

func anonFunction() {
	nextInt := intSeq()
	for i := 0; i < 5; i++ {
		fmt.Println(nextInt())
	}

	newInts := intSeq()
	fmt.Println(newInts())
}
