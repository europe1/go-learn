package main

import (
  "fmt"
  "time"
)

func lesson4() {
  const (
    Nanosecond time.Duration = 1
    Microsecond = 1000 * Nanosecond
    Millisecond = 1000 * Microsecond
    Second = 1000 * Millisecond
    Minute = 60 *  Second
    Hour = 60 * Minute
  )

  now := time.Now()
  var fiveMinutes time.Duration = 5 * Minute
  now = now.Add(-fiveMinutes)
  fmt.Println(now.Minute())

  fmt.Println("\niota")
  const (
    a1 = iota
    a2
    b1
    b2
  )
  fmt.Println(a1, a2, b1, b2)
}
