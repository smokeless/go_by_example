package main

import "fmt"

func vars() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2 //this is a really ugly way to do this.
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	var e int //not initialised
	fmt.Println(e)
	f := "short" //let the compiler infer it.
	fmt.Println(f)
}
