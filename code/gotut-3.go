package main

import (
	"fmt"
)

func add(x, y float64) float64 {
	return x + y
}
func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	// num1, num2 := 5.6, 9.5
	w1, w2 := "Hey", "there"
	fmt.Println(multiple(w1, w2))
}
