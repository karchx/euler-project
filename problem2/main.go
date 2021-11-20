package main

import "fmt"

var a []int

func fib() func(int) int {
	return func(x int) int {

		if x >= 2 {
			xf := a[x-1] + a[x-2]
			a = append(a, xf)
			return xf
		}
		a = append(a, x)
		return (a[x])
	}
}

func main() {
	var limit int = 4000000
  var result int
	f := fib()
	for i := 0; i <= limit; i++ {
		if i%2 == 0 {
      result = f(i)
      fmt.Print(result)
		}
	}

}
