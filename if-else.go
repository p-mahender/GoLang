package main

import "fmt"

func main() {
	if 5%2 == 0 {
		fmt.Println("5 is even number")
	} else {
		fmt.Println("5 is Odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 7%2 == 0 || 8%2 == 0 {
		fmt.Println("Either 8 or 7 is even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "num has 1 digit")
	} else {
		fmt.Println(num, " num has more than 1 digit")
	}

}
