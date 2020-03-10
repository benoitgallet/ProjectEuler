package main

import (
	"fmt"
)

func main() {
	var toFind = 600851475143
	//var sqrtToFindTmp = math.Sqrt(float64(toFind))
	//var sqrtToFind = int(sqrtToFindTmp)
	var maxPrimeFactor = 0
	for i := 2; i <= toFind; i ++ {
		for 0 == toFind % i {
			if maxPrimeFactor < i {
				maxPrimeFactor = i
			}
			toFind /= i
		}
	}
	fmt.Println(maxPrimeFactor)
}
