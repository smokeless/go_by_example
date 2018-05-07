package main

import "fmt"

func goMaps() {
	m := make(map[string]int) //empty map str:int
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))
	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"] //optional second return to tell whether or not
	//key was present in the map.
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
