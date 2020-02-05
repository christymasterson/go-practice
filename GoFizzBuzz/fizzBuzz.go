package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(num int) string {
	switch {
	case (num%3 == 0) && (num%5 == 0):
		return "FizzBuzz"
	case (num%3 == 0):
		return "Fizz"
	case (num%5 == 0):
		return "Buzz"
	default:
		return strconv.Itoa(num)
	}
}

func main() {
	for i := 0; i <= 15; i++ {
		fmt.Println(fizzBuzz(i))
	}
}
