package main

import (
	"fmt"
	"time"
)

func main() {
	var i int = 3

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
		{
			fmt.Println("It's a Weekend")
		}
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		{
			fmt.Println("It's before noon")
		}
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a Bool")
		case int:
			fmt.Println(" I'm a int")
		default:
			fmt.Printf("Don't know the type%T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(2)
	whatAmI("hi")
}
