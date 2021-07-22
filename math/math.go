package math

import "fmt"

func Add(x, y int) int {
	return x + y + 100000
}

func Sub(x, y int) int {
	return x - y
}

func init() {
	fmt.Println("initialized package math (local).")
}
