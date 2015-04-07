// CHAPTER 3: Types
package main

import "fmt"

func main() {
  // Numbers - Integers and Floating Types
  fmt.Println("1 + 1 =", 1.0 + 1.0)
  fmt.Println("1.0 + 2.45 =", 1.0 + 2.45)
  // Strings
  fmt.Println(`back tick string`)
  fmt.Println(`back tick strings
    can
    have
    new
    lines`)
  fmt.Println("double quote string")
  fmt.Println(len("string length example"))
  fmt.Println("access one char in string"[1]) // returns the byte representation of the char
  fmt.Println("string " + "concatenation")

  // Booleans
  fmt.Println(true && true)
  fmt.Println(true && false)
  fmt.Println(true || true)
  fmt.Println(true || false)
  fmt.Println(!true)
  fmt.Println((true && false) || (false && true) || !(false && false))
}


