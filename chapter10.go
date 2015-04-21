// Chapter 10: Concurrency
package main

import (
  "fmt"
  "time"
  "math/rand"
)

func f(n int) {
  for i := 0; i <10; i++ {
    fmt.Println(n, ":", i)
    amt := time.Duration(rand.Intn(250))
    time.Sleep(time.Millisecond * amt)
  }
}

// Channels
func pinger(c chan string) {
  for i := 0; ; i ++ {
    c <- "ping"
  }
}

func printer(c chan string) {
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}

func ponger(c chan string) {
  for i := 0; ; i ++ {
    c <- "pong"
  }
}

func main() {
  // Goroutines
  // for i := 0; i < 10; i++ {
  //   go f(i)
  // }z
  // var input string
  // fmt.Scanln(&input)

  // Channels
  var c chan string = make(chan string)

  go pinger(c)
  go ponger(c)
  go printer(c)

  var input string
  fmt.Scanln(&input)

}
