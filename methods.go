package main

import "fmt"

// type rect struct {
// 	width, height int
// }

// func (r *rect) area() int { //receiver type of *rect
// 	return r.width * r.height
// }

// func (r rect) perim() int { //non pointer receiver.
// 	return 2*r.width + 2*r.height
// }

func goMethods() {
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
