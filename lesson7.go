package main

import (
  "fmt"
  "sort"
)

type User struct {
  shares int
}

func lesson7() {
  users := map[string]User {
    "Doug": {12},
    "Valerie": {1},
  }
  users["Rob"] = User{7}
  users["Samantha"] = User{23}

  for k, v := range users {
    fmt.Println(k, v)
  }
  fmt.Println()
  var names []string
  for k := range users {
    names = append(names, k)
    fmt.Println(k)
  }

  fmt.Println()
  sort.Strings(names)
  for i := range names {
    fmt.Println(names[i])
  }

  delete(users, "Valerie")
  u, found := users["Valerie"]
  fmt.Println("\nusers['Valerie']", u, found)
}
