package main

import "fmt"
import "errors"

//errors get a seperate return value in go.
//no more try: except: finally:

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42") //so helpful.
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string { //attach this method to our struct.
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) { //error is last value
	if arg == 42 {
		return -1, &argError{arg, "can't work with it."} //reference to arg error.
	}
	return arg + 3, nil
}

func goErrors() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)

	}
}
