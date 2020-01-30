package main

import "fmt"

func lesson3() {
  var human1 Human
  var human2 *Human
  human1 = createHuman1()
  human2 = createHuman2()
  fmt.Printf("%+v\n%+v\n", human1, human2)
}

type Human struct {
  height float32
  weight float32
  age int
}

func createHuman1() Human {
  var h Human
  h.height = 193.0
  h.weight = 95.1
  h.age = 35
  fmt.Println(&h)
  return h
}

func createHuman2() *Human {
  var h Human
  h.height = 193.0
  h.weight = 95.1
  h.age = 35
  fmt.Println(&h)
  return &h
}
