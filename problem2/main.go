package main

import "fmt"

func main() {
	limit := 4_000_000
	var aux, fib uint64 = 1, 0
	var result uint64

	for i := 1; i <= limit; i++ {
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
