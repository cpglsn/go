package main

import (
  "fmt"
  "os"
)

type menuItem struct {
  name string
  price map[string]float64
}

func main() {

  menu := []menuItem{
    {name: "primi", price: map[string]float64{ "pasta": 1.1, "brodo": 1.2}},
    {name: "secondi", price: map[string]float64{ "carne": 2.1, "patate": 2.2}},
  }

  fmt.Println("What do you want me to do?")
  fmt.Println("1) Print first dishes")
  fmt.Println("2) Print second dishes")
  in := bufio.NewReader(os.Stdin)
  s, _ := in.ReadString('\n')

  for _, v := range menu {
    fmt.Println(v.name)
    for i2, v2 := range v.price {
      fmt.Println( "Piatto:", i2, "Prezzo:", v2 )
    }
  }
}

