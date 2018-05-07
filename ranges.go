package main

import "fmt"

func goRange() {
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	//range provides index and value, _ index, num value.

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "bear"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("keys:", k)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}
	/*
		range on strings iterates over Unicode code points.
		The first value is the starting byte index of the
		rune and the second the rune itself.
	*/
}
