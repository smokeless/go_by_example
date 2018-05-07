package main

import "fmt"

func vals() (int, int) { //(int, int) shows return 2 ints.
	return 3, 7
}

func multipleReturnValues() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
