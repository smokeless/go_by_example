package main

import "fmt"

func goSlice() {
	//slices are more powerful arrays.
	s := make([]string, 3) //slices are typed by the elements they contain
	s[0] = "whoa"
	s[1] = "another"
	s[2] = "array type"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "need another element")
	s = append(s, "and another", "and another")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s) //copy s into c
	fmt.Println("cpy:", c)

	l := s[2:5] //2,3,4
	fmt.Println("sl1:", l)

	l = s[:5] //up to but excludes 5

	l = s[2:] //up to and including 2

	t := []string{"g", "h", "i"} //declare slice on single line.
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
