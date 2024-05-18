package main

import (
	"fmt"
	"math"
)

func calc(x float64, y float64) float64 {
	// 2 乗の和の「平方根」を返す
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

func main() {
	fmt.Println(calc(3.0, 4.0)) // 5
}
