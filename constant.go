package main

import (
	"fmt"
	"math"
)

const C string = "constant"

func main() {
	fmt.Println(C)

	const n = 50000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
