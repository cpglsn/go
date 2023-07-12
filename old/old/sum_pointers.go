package main

import (
    "fmt"
    "time"
)


func sqr_a(numbers *[]int) int {
  var tmp int
	for item := 0; item < len(*numbers) ; item++ {
		tmp += (*numbers)[item]+1
	}
	return tmp
}

func sqr_b(numbers *[]int) int {
  var tmp int
	for item := 0; item < len(*numbers) ; item++ {
		tmp += (*numbers)[item]+2
	}
	return tmp
}

func sqr_c(numbers *[]int) int {
  var tmp int
	for item := 0; item < len(*numbers) ; item++ {
		tmp += (*numbers)[item]+3
	}
	return tmp
}


func main() {

  var lunghezza int = 1000000000

  numbers := make([]int, lunghezza)

  for item := 0; item < lunghezza ; item++ {
		numbers[item] = item
	}

  before_s := time.Now()

  a := sqr_a(&numbers)
  b := sqr_b(&numbers)
  c := sqr_c(&numbers)

  after_s := time.Now()

  fmt.Println(before_s, "\n", after_s, "\n", a, "\n", b, "\n", c)

}

