package main

import (
    "strconv"
)

func FizzBuzz(a int) string {
    var result = ""
	
	if a%5 == 0 && a%3 == 0 {
		result = "fizzbuzz"
	} else if a%3 == 0 {
		result = "fizz"
	} else if a%5 == 0 {
		result = "buzz"
	} else {
		result = strconv.Itoa(a)
	}
	
	return result
}
