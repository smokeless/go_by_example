package main

import "fmt"

//args passed to this function are passed by value.
//ie it's a copy.
func zeroval(ival int) {
	ival = 0
}

//args passed to this function are passed by reference.
func zeroptr(iptr *int) {
	*iptr = 0
}

func pointers() {
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i) //should be 1
	zeroptr(&i)                //pass by reference
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
}
