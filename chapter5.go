// CHAPTER 5: Control Structures
package main

import "fmt"

func main() {
  i := 1

  // For
  for i <= 10 {
    fmt.Println(i)
    i += 1
  }

  for j := 1; j <= 10; j++ {
    fmt.Println(j)
  }

  for j := 10; j >= 0; j-- {
    fmt.Println(j)
  }

  // If, else, else if
  for x := 1; x <= 10; x++ {
    if x % 2 == 0 {
      fmt.Println(x, "even")
    } else {
      fmt.Println(x, "odd")
    }
  }

  // getEnglishName()
  // printDivisibleByThree()
  printFizzDivisibleByThreePrintFiveDivisibleByFive()
}


// Switch
func getEnglishName() {
  num := 5
  switch num {
    case 0: fmt.Println("Zero")
    case 1: fmt.Println("One")
    case 2: fmt.Println("Two")
    case 3: fmt.Println("Three")
    case 4: fmt.Println("Four")
    case 5: fmt.Println("Five")
    default: fmt.Println("Unknown Number")
  }
}


func printDivisibleByThree() {
  for i := 1; i <= 100; i++ {
    if i % 3 == 0 {
      fmt.Println(i)
    }
  }
}

func printFizzDivisibleByThreePrintFiveDivisibleByFive() {
  for i := 1; i <= 100; i ++{
    if i % 3 == 0 {
      fmt.Println("Fizz")
    } else if i % 5 == 0 {
      fmt.Println("Buzz")
    } else {
      fmt.Println(i)
    }
  }
}
