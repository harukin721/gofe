package main

import "fmt"

func fizzBuzz(num int) string {
	var result string
	if num%3 == 0 && num%5 == 0 {
		result = "3と5で割り切れる"
	} else if num%3 == 0 {
		result = "3で割り切れる"
	} else if num%5 == 0 {
		result = "5で割り切れる"
	} else {
		result = "3でも5でも割り切れない"
	}
	return result
}

func main() {
	fmt.Println(fizzBuzz(15))
	fmt.Println(fizzBuzz(9))
	fmt.Println(fizzBuzz(10))
	fmt.Println(fizzBuzz(7))
}
