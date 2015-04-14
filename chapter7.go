package main

import "fmt"


func main() {
  xs := []float64{98,93,77,82,83}
  intSlice := []int{4,5,6}
  fmt.Println(average(xs))
  fmt.Println(sum(1,2,3))
  fmt.Println(sum(intSlice...))

  // Closure: functions inside of functions
  addTwo := func(x, y int) int {
    return x + y
  }
  fmt.Println(addTwo(10,10))

  nextEven := makeEvenGenerator()
  fmt.Println(nextEven()) // 0
  fmt.Println(nextEven()) // 2
  fmt.Println(nextEven()) // 4

  fmt.Println(factorial(5))

  // Defer, Panic & Recover
  // Similar to Python's try-catch
  // defer func() {
  //   str := recover()
  //   fmt.Println(str)
  // }()
  // panic("PANIC")

  // Exercises
  fmt.Println("PROBLEMS")
  fmt.Println("Sum is: ", sum(5,5,5))

  half, isEven := checkIfHalfIsEven(20)
  fmt.Println("Half is: ", half)
  fmt.Println("Even?", isEven)

  fmt.Println("Greatest number is: ", getGreatest(99, 5, 18, 81))

  oddGen := makeOddGenerator()

  fmt.Println(oddGen())
  fmt.Println(oddGen())
  fmt.Println(oddGen())
  fmt.Println(oddGen())

  fmt.Println(fibonacci(0), fibonacci(1), fibonacci(2), fibonacci(3), fibonacci(4), fibonacci(5))
}


func average(xs []float64) float64 {
  total := 0.0
  for _, v := range xs {
    total += v
  }
  return total / float64(len(xs))
}


// `sum` is a function which takes a slice of numbers and adds them together.
// This is how it will look like in Go.
func sum(args ...int) int {
  // Similar to kwargs in Python
  total := 0
  for _, v := range args {
    total += v
  }
  return total
}

// A function which takes an integer and halves it and returns true
// if it was even or false if it was odd. For example, `half(1)` should
// return `(0, false)` and `half(2)` should return `(1, true)`.
func checkIfHalfIsEven(num int) (int, bool) {
  half := num / 2

  if half % 2 == 0 {
    return half, true
  } else {
    return half, false
  }
}

// Similar to Python generators
// Function that generates even numbers
func makeEvenGenerator() func() uint {
  i := uint(0)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}


// Function that generates odd numbers
func makeOddGenerator() func() uint {
  i := uint(1)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}


// Recursion
func factorial(x uint) uint {
  if x == 0 {
    return 1
  }
  return x * factorial(x-1)
}


// Function that finds the greatest number in a list of numbers
func getGreatest(args ...int) int {
  greatest := 0
  for i, v := range args {
    if i == 0 {
      greatest = v
    } else {
      if v > greatest {
        greatest = v
      }
    }
  }

  return greatest
}

func fibonacci(num int) (int) {
  if num == 0 {
    return 0
  } else if num == 1 {
    return 1
  } else {
    return fibonacci(num-1) + fibonacci(num-2)
  }
}
