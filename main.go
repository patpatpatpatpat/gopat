// Source: http://www.golang-book.com/

// CHAPTER 2: Your First Program

package main

import "fmt"
import "os"

// welcome to the go world, pat!

/*
Another way of writing comments in Go
*/

func main(){
  fmt.Println("Hello, my name is Pat!")
  os.Exit(0)
}
