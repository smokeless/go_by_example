package main

import "fmt"

func forLoops() {
	//for is the only looping construct in go?
	i := 1
	for i <= 3 { //this may as well be while.
		fmt.Println(i)
		i++
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
