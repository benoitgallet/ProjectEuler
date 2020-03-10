package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	var tmp = string(n)
	var itBegin = 0
	var itEnd = len(tmp) - 1
	for itBegin < itEnd {
		if tmp[itBegin] != tmp[itEnd] {
			return false
		}
		itBegin++
		itEnd--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(9009))
	var nb1, nb2 int
	var maxPalindrome = 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			var tmp = i * j
			if isPalindrome(tmp) {
				if maxPalindrome < tmp {
					nb1 = i
					nb2 = j
				}
			}
		}
	}
	fmt.Println(strconv.Itoa(nb1) + ", " + strconv.Itoa(nb2))
}
