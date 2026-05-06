package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	root := float64(1)
	for i := 0; i < 10; i++ {
	  fmt.Printf("Root: %f\nIteration: %d\n", root, i)
	  root -= (root*root - x) / (2*root)
	}
	
	return root
}

func main() {
	Sqrt(2)
}

