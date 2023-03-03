package main

// Works perfectly fine on the website but it can't find the module locally
// Probably some way to fix that, but i wont deal with it now.
// need more training time before i try that.

// But should work for you if you have the *thing* imported

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	array := make([][]uint8, 0, dy)
	for y := 0; y < dy; y++ {
		innerArray := make([]uint8, 0, dx)
		for x := 0; x < dx; x++ {
			innerArray = append(innerArray, uint8(mathFunc(x, y)))
		}
		array = append(array, innerArray)
	}

	return array
}

func mathFunc(x, y int) uint8 {
	return uint8((x ^ y))
}

func main() {
	pic.Show(Pic)
}
