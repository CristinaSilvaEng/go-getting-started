package main

import "fmt"

// constant block
const (
  first = 1 << (10 * iota)
  second
  third
)

func main() {
  // int type
  var response int
  response = 42
  println(response)

// float type
  var dolar float32
  dolar = 3.65
  println(dolar)

// string type
  goLang := "Go Language!"
  println(goLang)

// complex type
  complexo := complex(2, 3)
  println(complexo)
  println(real(complexo))
  println(imag(complexo))

// call constants
  println(first)
  println(second)
  println(third)

// array
  arrray := [3]int{}
  arrray[0] = 40
  arrray[1] = 41
  arrray[2] = 42
  fmt.Println(arrray)

// initialized array
  arrrray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  fmt.Println(arrrray)
  fmt.Println(len(arrrray))

// slice type
  slicee := arrrray[1:7]
  fmt.Println(slicee)

  slicee = arrrray[:]
  fmt.Println(slicee)

// slice add value
  slicee = append(slicee, 42)
  fmt.Println(slicee)

// slice initialized size
  sliceee := make([]float32, 100)
  sliceee[1] = 1.
  sliceee[2] = 2.
  sliceee[3] = 3.
  fmt.Println(sliceee)
  fmt.Println(len(sliceee))

// map type
  maap := make(map[int]string)
  fmt.Println(maap)
  maap[42] = "O universo, a vida e tudo mais..."
  maap[0] = "Foo"
  maap[1] = "Bar"

  fmt.Println(maap)
  fmt.Println(maap[99])

// operators
  add := 1 + 2
  println(add)

  subtract := 10 -5
  println(subtract)

  remainder := 5 % 2
  println(remainder)

  div := 5 / 2
  println(div)

  inc := 1
  inc++
  println(inc)
  inc++
  println(inc)
  inc += -3
  println(inc)
}
