package main

import (
	"fmt"
)

func main() {
	numbers := [3]float64{71.2, 56.3, 89.1}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum)
}
