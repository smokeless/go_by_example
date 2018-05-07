package main

import "fmt"

func fact(n int) int { //This fact function calls itself
	//until it reaches the base case of fact(0).
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func functionRecursion() {
	fmt.Println(fact(7))
}
