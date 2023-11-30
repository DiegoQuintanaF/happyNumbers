package main

import (
	"fmt"
	"math"

	"github.com/diegoquinfa/happyNumbers/hn"
)

func main() {
	var a int64 = 1
	var b int64 = int64(math.Pow(10, 16))
	test := hn.NewHappyNumber()
	result := test.CountHappyNumbers(a, b)

	fmt.Printf("%d\n", result)
}
