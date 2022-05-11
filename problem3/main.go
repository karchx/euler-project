package main

import (
	"fmt"
	"math"
)

func main() {
	n := largestPrimeFactor(600851475143)
	fmt.Println(n)
}

func largestPrimeFactor(number int64) int64 {
	var lastFactor, factor, maxFactor int64

	if math.Mod(float64(number), 2) == 0 {
		lastFactor = 2
		number = number / 2

		for i := 0; math.Mod(float64(number), 2) == 0; i++ {
			number = number / 2
		}
	} else {
		lastFactor = 1
	}
	factor = 3
	maxFactor = int64(math.Sqrt(float64(number)))

	for i := 0; number > 1 && factor <= int64(maxFactor); i++ {
		if math.Mod(float64(number), float64(factor)) == 0 {
			number = number / factor
			lastFactor = factor
			maxFactor = int64(math.Sqrt(float64(number)))
		}
		factor = factor + 2
	}

	if number == 1 {
		return lastFactor
	} else {
		return number
	}
}
