package main

import (
	"fmt"
)

func main() {
	fmt.Println(Fib(0))
	fmt.Println(Fib(1))
	fmt.Println(Fib(2))
	fmt.Println(Fib(3))
	fmt.Println(Fib(4))
	fmt.Println(Fib(5))
	fmt.Println(Fib(6))
	fmt.Println(Fib(7))
}

func Fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}
