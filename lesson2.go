package main

import "fmt"

func lesson2() {
  var m int
  m = 115

  fmt.Println("Passing a value")

  fmt.Printf("m = %d; the address is %v\n", m, &m)
  incr(m)
  fmt.Printf("m = %d; the address is %v\n", m, &m)

  fmt.Println()
  fmt.Println("Passing an address")

  fmt.Printf("m = %d; the address is %v\n", m, &m)
  incr_pointer(&m)
  fmt.Printf("m = %d; the address is %v\n", m, &m)
}

func incr(a int) {
  fmt.Printf("a = %d; the address is %v\n", a, &a)
  a++
}

func incr_pointer(a *int) {
  fmt.Printf("a = %d; the address is %v\n", a, &a)
  fmt.Printf("*a = %d; the address is %v\n", *a, &*a)
  *a++
}
