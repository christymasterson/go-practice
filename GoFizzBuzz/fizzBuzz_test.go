package main

import (
	"testing"
)

func TestFizzBuzz_FizzBuzz(t *testing.T) {

	if fizzBuzz(15) != "FizzBuzz" {
		t.Error("Expected returned string to be 'FizzBuzz'")
	}
}
func TestFizzBuzz_Fizz(t *testing.T) {

	if fizzBuzz(3) != "Fizz" {
		t.Error("Expected returned string to be 'Fizz'")
	}
}

func TestFizzBuzz_Buzz(t *testing.T) {

	if fizzBuzz(5) != "Buzz" {
		t.Error("Expected returned string to be 'Buzz'")
	}
}
