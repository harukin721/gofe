package main

import "fmt"

func binSort(data []int) []int {
	// n は data の要素数
	n := len(data)
	// bins は data の最大値までの要素数を持つ配列
	bins := make([]int, n+1)
	// data の要素を bins に格納
	for i := 0; i < n; i++ {
		bins[data[i]] = data[i]
	}
	// bins から 0 以外の要素を取り出して result に格納
	result := []int{}
	for i := 1; i <= n; i++ {
		if bins[i] != 0 {
			// 0 以外の要素を result に追加
			result = append(result, bins[i])
		}
	}
	return result
}

func main() {
	// data はソート対象の配列
	data := []int{2, 6, 3, 1, 4, 5}

	// 以下は不正解の解答群
	// data := []int{3, 1, 4, 4, 5, 2}
	// data := []int{4, 2, 1, 5, 6, 2}
	// data := []int{5, 3, 4, 3, 2, 6}
	fmt.Println(binSort(data))
}
