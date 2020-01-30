package main

import "fmt"

func lesson1() {
  type A struct {
    name string
    active bool
  }

  type B struct {
    name string
    active bool
  }

  var stefan A
  var claudia B
  claudia.name = "Claudia"
  stefan = A(claudia)
  fmt.Printf("%+v\n", stefan)
}
