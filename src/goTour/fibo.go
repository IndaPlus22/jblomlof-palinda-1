package main

import "fmt"

func fibonacci() func() int {
	var current, next int = 0, 1

	return func() int {
		res := current
		current = next
		next = res + next
		return res
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
