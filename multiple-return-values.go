package main

import "fmt"

func val(a int, b int) (int, int) {
	return a + b, a - b
}

func main() {
	sum, diff := val(3, 4)
	fmt.Println("3+4:", sum, "3-4:", diff)

	_, dif := val(5, 6)
	fmt.Println("dif:", dif)
}
