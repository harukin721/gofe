package main

import "fmt"

// Go には while 文がないので for 文で代用
func gcd(num1 int, num2 int) int {
	x := num1
	y := num2
	for x != y {
		if x > y {
			x = x - y
		} else {
			y = y - x
		}
	}
	return x
}

func main() {
	fmt.Println(gcd(5, 5))     // 5
	fmt.Println(gcd(48, 18))   // 6
	fmt.Println(gcd(101, 103)) // 1
	fmt.Println(gcd(270, 192)) // 6
}
