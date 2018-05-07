package main

import "fmt"
import "time"

func goSwitch() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend.")
	default:
		fmt.Println("it's a weekday.")
	}

	t := time.Now()

	switch { //different if else
	case t.Hour() < 12:
		fmt.Println("It's before noon.")
	default:
		fmt.Println("It's afternoon.")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("i'm an int.")
		default:
			fmt.Println("Unknown type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
