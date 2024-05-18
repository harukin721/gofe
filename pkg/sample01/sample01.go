package main

import "fmt"

func main() {
	var x int = 1
	var y int = 2
	var z int = 3

	x = y
	y = z
	z = x

	fmt.Printf("%d,%d\n", y, z)
}
