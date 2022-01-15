package main

import "fmt"

func primeFactor() func(int)int{
  return func (n int)int {
    return n
  }
}

func main() {
  pf := primeFactor()
  fmt.Print(pf(2))
}
