// Chapter 8: Pointers
package main

import "fmt"

func main() {
  x := 5
  zero(&x)
  fmt.Println(x)

  y := new(int)
  one(y)
  fmt.Println(*y)

  // Exercises
  // Program that can swap two integers `x := 1; y:= 2; swap(&x, &y)` should give you `x=2` and `y=1`
  a := 1
  b := 2
  swap(&a, &b)
  fmt.Println(a)
  fmt.Println(b)
}

func zero(num *int) {
  *num = 0
}

func one(num *int) {
  *num = 1
}

func swap(first *int, second *int) {
  *first, *second = *second, *first
}
