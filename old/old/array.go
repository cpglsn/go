package main

import "fmt"

func cambia(a [3]int) {
  a[0] = 1
  fmt.Println(a)
}

func main() {
  b := [3]int{4,5,6}
  cambia(b)
  fmt.Println(b)
}

