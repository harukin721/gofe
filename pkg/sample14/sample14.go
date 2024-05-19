package main

import (
	"fmt"
	"log"
	"math"
)

// 要素数が 1 以上で、昇順に整列済みの配列を基に、配列を特徴づける 5 つの値を求める
func findRank(sortedData []float64, p float64) float64 {
	i := int(math.Ceil(p * float64(len(sortedData)-1)))
	return sortedData[i]
}

func summarize(sortedData []float64) []float64 {
	// 要素数 0 の配列
	rankData := []float64{}
	p := []float64{0, 0.25, 0.5, 0.75, 1}

	// p の各要素に対して、findRank 関数を適用
	for _, v := range p {
		log.Printf("v: %f\n", v)
		// rankData に findRank 関数の結果を追加
		rankData = append(rankData, findRank(sortedData, v))
		log.Printf("rankData: %v\n", rankData)
	}

	return rankData
}

func main() {
	data := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1}
	fmt.Println(summarize(data)) // [0.1 0.4 0.6 0.8 1]
}

// 出力
// v: 0.000000
// rankData: [0.100000]
// v: 0.250000
// rankData: [0.100000 0.400000]
// v: 0.500000
// rankData: [0.100000 0.400000 0.600000]
// v: 0.750000
// rankData: [0.100000 0.400000 0.600000 0.800000]
// v: 1.000000
// rankData: [0.100000 0.400000 0.600000 0.800000 1.000000]
// [0.1 0.4 0.6 0.8 1]
