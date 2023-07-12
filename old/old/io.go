package main

import (
  "fmt"
  "strings"
  "os"
  "bufio"
)

func main() {
  fmt.Println("What do you want me to scream?")
  in := bufio.NewReader(os.Stdin)
  s, _ := in.ReadString('\n')
  fmt.Println(strings.ToUpper(strings.TrimSpace(s)) + "!")
}

