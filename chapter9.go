// Chapter 9: Structs and Interfaces
package main

import ("fmt"; "math")

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 -y1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}


// Structs
type Circle struct {
  x, y, r float64
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}

// Methods
func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}

func (c *Circle) perimeter() float64 {
  return 2*math.Pi*c.r
}


func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}

func (r *Rectangle) perimeter() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return 2 * (l+w)
}

// Embedded Types
type Person struct {
  Name string
}

func (p *Person) Talk() {
  fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
  Person
  Model string
}


// Interfaces
type Shape interface {
  area() float64
  perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
  var area float64
  for _, s := range shapes {
    area += s.area()
  }
  return area
}

type MultiShape struct {
  shapes []Shape
}

func (m *MultiShape) area() float64 {
  var area float64
  for _, s := range m.shapes {
    area += s.area()
  }
  return area
}


func main() {
  // Initialization
  // var c Circle
  // c := new(Circle)

  // Initialization with values
  c := Circle{x: 0, y: 0, r: 5}
  // c := Circle{0, 0, 5}

  fmt.Println(c)
  fmt.Println(c.x, c.y, c.r)
  // fmt.Println(circleArea(c))
  // fmt.Println(circleArea(&c)) // Use if one of the fields will be modified
  fmt.Println("Circle area is:", c.area())
  fmt.Println("Circle perimeter is:", c.perimeter())


  r := Rectangle{0, 0, 10, 10}
  fmt.Println(r)
  fmt.Println("Rectangle area is:", r.area())
  fmt.Println("Rectangle perimeter is:", r.perimeter())

  // Embedded Types
  a := Android{Model: "v1.0"}
  a.Person.Talk()
  a.Talk()

  // Interfaces
  fmt.Println(totalArea(&c, &r))
}




