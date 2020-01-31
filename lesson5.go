package main

import "fmt"

func lesson5() {
  var names1 = []string{"John", "Mike", "Sarah", "Swenn"}
  var names2 []string

  names2 = names1[1:3]
  fmt.Println("\nnames1")
  printSlice(names1)
  fmt.Println("names2")
  printSlice(names2)

  names2[0] = "Isaiah"
  names2 = append(names2, "Julia")
  fmt.Println()
  fmt.Println("\nnames1")
  printSlice(names1)
  fmt.Println("names2")
  printSlice(names2)

  var names3 []string
  names3 = names1[1:3:3]
  names3[0] = "Maxwell"
  names3 = append(names3, "Antony")
  fmt.Println("\nnames1")
  printSlice(names1)
  fmt.Println("names3")
  printSlice(names3)

  names4 := make([]string, len(names1))
  copy(names4, names1)
  fmt.Println("\nnames1")
  printSlice(names1)
  fmt.Println("names4")
  printSlice(names4)
}

func printSlice(slice []string) {
  fmt.Printf("length: %d, capacity: %d\n", len(slice), cap(slice))
  for i, v := range slice {
    fmt.Printf("%d | %v | \"%s\"\n", i, &slice[i], v)
  }
}
