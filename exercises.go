package main

import (
	"fmt"
)

func loopTenTimes() int {
	var a int
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		a = i
	}
	return a
}

func fizzBuzz() {
	for i := 1; i < 101; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
