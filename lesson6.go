package main

import (
  "fmt"
  "unicode/utf8"
)

func lesson6() {
  var src string = "That was a 漫长的一天"
  var buf [utf8.UTFMax]byte
  for i, r := range src {
    rl := utf8.RuneLen(r)
    si := i + rl
    copy(buf[:], src[i:si])
    fmt.Printf("%02d: %c %#v\n", i, r, buf[:rl])
  }
}
