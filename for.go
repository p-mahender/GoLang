package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println("i:", i)
		i = i + 1
	}

	for j := 4; j <= 7; j++ {
		fmt.Println("j:", j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("n:", n)
	}

}
