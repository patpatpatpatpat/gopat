// CHAPTER 4:
// Variables, Scopes, Constants
// Defining Multiple Variables
package main

import "fmt"

var adj string = "Awesome"

// Trivia Nights Constants
const winCount int = 8
const eventLocation string = "Sales Tekanplor"
const eventSetCount int = 4
const categoriesPerSetCount int = 3

// More Trivia Nights Constants
const (
  defaultGameMaster = "Chris"
  defaultFirstPrize = "Crispy Pata"
  defaultSecondPrize = "1 Bucket of Beer"
)

func main() {
  var x string = "Hello world"
  var name string
  name = "Pat"
  fmt.Println(name)

  x = "first"
  fmt.Println(x)
  x = "second"
  fmt.Println(x)

  x = "hello"
  var y string = "world"
  fmt.Println(x == y)


  myNum := 1
  myStr := "hi"
  myFloat := 0.138

  fmt.Println(myNum)
  fmt.Println(myStr)
  fmt.Println(myFloat)

  printAdjective()
  getWord()

  fmt.Println("Number of Trivia Night wins: ", winCount)

  // doubleYourNumber()
  // fahrenheitToCelsius()
  // feetToMeters()
}


func getWord() {
  fmt.Println("Printing your word: ", adj)
}

func printAdjective() {
  fmt.Println("Printing your adjective: ", adj)
}

func doubleYourNumber() {
  /*
  This function takes in a number
  entered by the user and doubles it.
  */
  fmt.Print("Enter a number: ")
  var input float64
  fmt.Scanf("%f", &input)

  output := input * 2

  fmt.Println("Your number doubled is: ", output)
}

func fahrenheitToCelsius(){
  /*
  This function takes in a Fahrenheit measure
  and returns the equivalent in Celsius
  */
  fmt.Print("Enter temperature in Fahrenheit: ")
  var input float64
  fmt.Scanf("%f", &input)

  celsius := (input - 32) * 5/9
  fmt.Println("Fahrenheit is: ", input)
  fmt.Println("Celsius is: ", celsius)
}

func feetToMeters(){
  /*
  This function takes in a measurement in Feet
  and returns the equivalent in Meters
  */
  fmt.Print("Enter measurement in Feet: ")
  var feetInput float64
  fmt.Scanf("%f", &feetInput)

  meter := feetInput * 0.3048
  fmt.Println("Fahrenheit is: ", feetInput)
  fmt.Println("meter is: ", meter)
}
