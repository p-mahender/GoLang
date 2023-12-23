package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 10

	fmt.Println("m:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("m:", m)

	clear(m)
	fmt.Println("m:", m)

	_, pre := m["k2"]
	fmt.Println("pre:", pre)

	n := map[string]int{"k1": 1, "k2": 2}
	fmt.Println("n", n)

	n2 := map[string]int{"k1": 1, "k2": 2}

	if maps.Equal(n, n2) {
		fmt.Println("n==n2")

	}

}
