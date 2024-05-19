package main

import (
	"fmt"
	"log"
)

// Unicode の符号位置を、UTF-8 の符号に変更する
func encode(codePoint int) []int {
	// utf8Bytes の初期値は，ビットパターンの "x" を全て 0 に置き換え、
	// 8桁ごとに区切って，それぞれを2進数とみなしたときの値
	utf8Bytes := []int{224, 128, 128}
	cp := codePoint

	// 6ビットずつに分割して，それぞれの値をutf8Bytesに格納
	for i := len(utf8Bytes) - 1; i >= 0; i-- {
		log.Printf("utf8Bytes[%d]: %d", i, utf8Bytes[i])
		// 6ビットずつに分割
		utf8Bytes[i] += cp % 64
		log.Printf("utf8Bytes[%d]: %d", i, utf8Bytes[i])
		// 6ビットずつに分割した値をutf8Bytesに格納
		cp /= 64
	}

	return utf8Bytes
}

func main() {
	// "あ" のUnicode符号位置をUTF-8の符号に変換
	fmt.Println(encode(0x3042)) // [227 129 130]
}

// 出力
// utf8Bytes[2]: 128
// utf8Bytes[2]: 130
// utf8Bytes[1]: 128
// utf8Bytes[1]: 129
// utf8Bytes[0]: 224
// utf8Bytes[0]: 227
// [227 129 130]
