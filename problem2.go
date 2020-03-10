package main

import "fmt"

func main() {
	var sum = 0
	var fib1 = 1
	var fib2 = 1
	var fibTmp = 0
	for fibTmp < 4000000 {
		fibTmp = fib1 + fib2
		if 0 == fibTmp % 2 {
			sum += fibTmp
		}
		fib1 = fib2
		fib2 = fibTmp
	}
	fmt.Println(sum)
}
