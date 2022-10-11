package main

import (
	"fmt"
)

func is_divided_by_3(number int) bool {
	return number%3 == 0
}

func is_divided_by_5(number int) bool {
	return number%5 == 0
}

func is_divided_by_15(number int) bool {
	return is_divided_by_3(number) && is_divided_by_5(number)
}

func main() {
	for i := 1; i <= 30; i++ {
		if is_divided_by_15(i) {
			fmt.Println("FizzBuzz")
		} else if is_divided_by_3(i) {
			fmt.Println("Fizz")
		} else if is_divided_by_5(i) {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
