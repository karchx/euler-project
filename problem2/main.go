package main

import "fmt"

func main() {
	var limit int = 4_000_000
	var aux, fib uint64 = 1, 0
	var init int
	var result uint64

	for init = 1; init <= limit; init++ {
		aux += fib
		fib = aux - fib
		if fib <= 4_000_000 {
			if fib%2 == 0 {
				result += fib
			}
		}
	}
	fmt.Println(result)
}
