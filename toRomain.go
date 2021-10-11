package main

import (
  "fmt"
)

func intToRoman(num int) string {
  symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
  values := []int {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
  var ans string
  for i := 0; num != 0; i++ {
    for ; num >= values[i]; {
      num -= values[i]
      ans += symbols[i]
    }
  }
  return ans
}

func main() {
  var num int
  fmt.Println("Write a number: ")
  fmt.Scan(&num)
  fmt.Println(intToRoman(num))
}
