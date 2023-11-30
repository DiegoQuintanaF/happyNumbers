package main

import (
	"fmt"

	"github.com/diegoquinfa/happyNumbers/hn"
)

func main() {
	var a int64
	var b int64

	fmt.Scanf("%d %d", &a, &b)
	test := hn.NewHappyNumber()
	result := test.CountHappyNumbers(a, b)

	fmt.Printf("%d\n", result)
}
