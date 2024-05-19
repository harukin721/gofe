package main

import (
	"fmt"
	"log"
)

// 引数 data で指定された配列に、引数 target で指定された値が含まれていればその要素番号を返し、含まれていなければ -1 を返す
func search(data []int, target int) int {
	low := 1
	high := len(data)
	log.Printf("low: %d, high: %d\n", low, high)

	// Go には while 文がないので for で代用
	for low <= high {
		log.Printf("low: %d <= high: %d\n", low, high)
		middle := (low + high) / 2

		if data[middle-1] < target {
			log.Printf("data[%d]: %d < target: %d\n", middle-1, data[middle-1], target)
			low = middle
		} else if data[middle-1] > target {
			log.Printf("data[%d]: %d > target: %d\n", middle-1, data[middle-1], target)
			high = middle
		} else {
			log.Printf("data[%d]: %d == target: %d\n", middle-1, data[middle-1], target)
			return middle
		}
	}
	return -1
}

func main() {
	// 「要素数が 2 で，target が data の先頭要素の値と等しい」場合にに無限ループにならない
	// 「要素数が 2 で，target が data の末尾要素の値と等しい」場合に無限ループになる
	data := []int{3, 8}
	target := 8
	fmt.Println(search(data, target))
}

// 出力
// low: 1, high: 2
// low: 1 <= high: 2
// data[0]: 3 < target: 8
// low:1 <= high: 1
// data[0]: 3 < target: 8
// low:1 <= high: 1
// data[0]: 3 < target: 8
// ...
