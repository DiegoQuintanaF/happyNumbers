package main

import (
	"fmt"
	"math"
	"time"

	"github.com/diegoquinfa/happyNumbers/hn"
)

func main() {
	var a int64 = 1
	var b int64 = int64(math.Pow(10, 16))
	test := hn.NewHappyNumber()
	start := time.Now()
	result := test.CountHappyNumbers(a, b)

	fmt.Printf("hay %d entre %d a %d y tardo %v\n", result, a, b, time.Since(start))
}
