package main

import "fmt"

// 単方向リストの要素を表す構造体
type ListElement struct {
	val  rune         // ノードの値
	next *ListElement // 次のノードへのポインタ
}

// リストの先頭要素を表す変数
var listHead *ListElement

// リストの要素を削除する関数
func delNode(pos int) {
	if pos == 1 {
		listHead = listHead.next
	} else {
		prev := listHead
		// pos が2と等しいときは繰り返し処理を実行しない
		for i := 2; i < pos; i++ {
			prev = prev.next
		}
		prev.next = prev.next.next
	}
}

func main() {
	// a -> b -> c のリストを作成
	listHead = &ListElement{val: 'a', next: &ListElement{val: 'b', next: &ListElement{val: 'c', next: nil}}}

	// リストの2番目の要素を削除
	delNode(2)

	// リストのすべての要素を出力
	for e := listHead; e != nil; e = e.next {
		fmt.Println(string(e.val))
	}
}
