package main

import (
	"fmt"
)

func main() {

	defer fmt.Println("one")
	defer fmt.Println("Two")
	fmt.Println("Three")

}
