// CHAPTER 6: Arrays, Slices and Maps
/*
  Arrays - numbered sequence of elements of a single type with a fixed length
  Slices - Segment of an array. Indexable and have a length. Length is allowed to change
*/
package main

import "fmt"

func main() {
  // Arrays
  var x [5]int
  x[4] = 100
  y := [5]float64{
    99,
    98,
    // 97,

    96,
    95,
  }

  fmt.Println(y)

  calculateTestScores()

  // Slices
  // var q []float64
  w := make([]float64, 5)  // Length: 5
  e := make([]float64, 5, 10) // Capacity: 10
  r := [5]float64{1, 2, 3, 4, 5}

  fmt.Println(w)
  fmt.Println(e)
  fmt.Println(r[0:])
  fmt.Println(r[0:len(r)])
  fmt.Println(r[0:5])
  fmt.Println(r[:])

  // Slice Functions: append
  slice1 := []int{1,2,3}
  slice2 := append(slice1, 4, 5)
  fmt.Println(slice2)

  // Slice Functions: copy
  slice3 := []int{5,6,7}
  slice4 := make([]int, 2)
  copy(slice3, slice4)
  fmt.Println(slice3, slice4)


  // Maps
  u := make(map[string]int)
  u["key"] = 10

  fmt.Println(u)


  i := make(map[int]int)
  i[1] = 10
  i[0] = 12
  fmt.Println(i)

  delete(i, 0)
  fmt.Println(i)

  elements := make(map[string]string)
  elements["H"] = "Hydrogen"
  elements["He"] = "Helium"
  elements["Li"] = "Lithium"
  elements["Be"] = "Beryllium"
  elements["B"] = "Boron"
  elements["C"] = "Carbon"
  elements["N"] = "Nitrogen"
  elements["O"] = "Oxygen"
  elements["F"] = "Fluorine"
  elements["Ne"] = "Neon"

  fmt.Println(elements["Li"])
  fmt.Println(elements["xxx"])
  name, ok := elements["xx"]
  fmt.Println(name, ok)

  elementz := map[string]map[string]string{
    "H": map[string]string{
        "name":"Hydrogen",
        "state":"gas",
    },
    "He": map[string]string{
        "name":"Helium",
        "state":"gas",
    },
    "Li": map[string]string{
        "name": "Lithium",
        "state": "solid",
      },
    "Be": map[string]string{
        "name":"Beryllium",
        "state":"solid",
    },
    "B": map[string]string{
        "name":"Boron",
        "state":"solid",
    },
    "C": map[string]string{
        "name": "Carbon",
        "state": "solid",
      },
    "N": map[string]string{
        "name": "Nitrogen",
        "state": "gas",
      },
    "O": map[string]string{
        "name":"Oxygen",
        "state":"Gas",
    },
    "F": map[string]string{
        "name":"Flourine",
        "state":"gas",
    },
    "Ne": map[string]string{
        "name":"Neon",
        "state":"gas",
    },
  }

  if el, ok := elementz["Li"]; ok {
    fmt.Println(el["name"], el["state"])
  }

  // Program that finds the smallest number in the list
  myList := []int{
    48, 96, 86, 68,
    57, 82, 63, 70,
    37, 34, 83, 27,
    19, 97, 9, 17,
  }
  var smallest int
  for i := 0; i < len(myList);  i++ {
    if i == 0 {
      smallest = myList[i]
    } else {
      if myList[i] < smallest {
        smallest = myList[i]
      }
    }
  }
  fmt.Println("Smallest number in list: ", smallest)
}


func calculateTestScores(){
  var x [5]float64
  x[0] = 98
  x[1] = 93
  x[2] = 77
  x[3] = 82
  x[4] = 83

  var total float64 = 0
  // for i := 0; i < len(x); i++ {
  //   total += x[i]
  // }

  for _, value := range x {
    total += value
  }
  fmt.Println(total / float64(len(x)))
}

