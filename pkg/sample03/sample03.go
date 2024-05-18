package main

import "fmt"

func makeNewArray(in []int) []int {
	out := make([]int, len(in))       // 入力配列と同じ長さの新しい配列を作成
	out[0] = in[0]                    // 新しい配列の最初の要素は入力配列の最初の要素と同じ
	for i := 1; i < len(in); i++ {    // 入力配列の2番目の要素から最後の要素までループ
		out[i] = out[i-1] + in[i] // 新しい配列の現在の要素は、その前の要素と入力配列の現在の要素の和
	}
	return out // 新しい配列を返す
}

func main() {
	fmt.Println(makeNewArray([]int{3, 2, 1, 6, 5, 4}))
}
