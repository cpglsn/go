package main

import (
    "fmt"
    "time"
    "math"
)


func sqr_a(numbers []int) float64 {
  var tmp float64
	for item := 0; item < len(numbers) ; item++ {
		tmp += math.Sqrt(math.Pow(float64(numbers[item]+1), 3))
	}
	return tmp
}

func sqr_b(numbers []int) float64 {
  var tmp float64
	for item := 0; item < len(numbers) ; item++ {
		tmp += math.Sqrt(math.Pow(float64(numbers[item]+1), 4))
	}
	return tmp
}

func sqr_c(numbers []int) float64 {
  var tmp float64
	for item := 0; item < len(numbers) ; item++ {
		tmp += math.Sqrt(math.Pow(float64(numbers[item]+1), 5))
	}
	return tmp
}


func main() {

  numbers := [100000000]int

  for item := 0; item < 100000000 ; item++ {
		numbers[item] = item
	}

  before_s := time.Now()

  var a = sqr_a(numbers)
  var b = sqr_b(numbers)
  var c = sqr_c(numbers)

  after_s := time.Now()

  fmt.Println(before_s, "\n", after_s, "\n", a, "\n", b, "\n", c)

}

