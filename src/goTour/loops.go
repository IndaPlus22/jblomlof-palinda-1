package main

import (
	"fmt"
	"math"
)

const allowedDiff = 1e-15

func sqrt(x float64) float64 {
	z := x / 2.0
	for diff := z*z - x; allowedDiff <= math.Abs(diff); diff = z * z  -x {
		z -= diff / (2*z)
	}
	return z
}

func main() {
	fmt.Println(sqrt(2), sqrt(3), sqrt(9), sqrt(0.25))
}
