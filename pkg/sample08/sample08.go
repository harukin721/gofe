package main

import (
	"container/heap"
	"fmt"
)

// Item は優先度付きキューの要素を表す
type Item struct {
	value    string // キューに入れる値
	priority int    // キューの優先度
	index    int    // キューのインデックス
}

// PriorityQueue は優先度付きキューを表す
type PriorityQueue []*Item

// Len は優先度付きキューの要素数を返す
func (pq PriorityQueue) Len() int { return len(pq) }

// Less は要素iが要素jよりも優先度が高いかどうかを返す
// 小さい値が高い優先度を意味する（min-heap）
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

// Swap は要素iと要素jを入れ替える
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push は要素xを追加する
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop は最小値を取り出す
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // メモリリークを避ける
	item.index = -1 // 安全のため
	*pq = old[0 : n-1]
	return item
}

// update はキュー内のアイテムの値と優先度を更新し、その位置を修正する
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// PrioQueue は PriorityQueue を使いやすくするためのラッパー
type PrioQueue struct {
	pq PriorityQueue
}

// NewPrioQueue は優先度付きキューを初期化して返す
func NewPrioQueue() *PrioQueue {
	pq := &PriorityQueue{}
	heap.Init(pq)
	return &PrioQueue{pq: *pq}
}

// Enqueue は値と優先度を指定してキューに追加する
func (pq *PrioQueue) Enqueue(value string, priority int) {
	item := &Item{
		value:    value,
		priority: priority,
	}
	heap.Push(&pq.pq, item)
}

// Dequeue は優先度が最も高い（最小の）値を取り出して返す
func (pq *PrioQueue) Dequeue() string {
	if pq.Size() == 0 {
		return ""
	}
	item := heap.Pop(&pq.pq).(*Item)
	return item.value
}

// Size はキューの要素数を返す
func (pq *PrioQueue) Size() int {
	return pq.pq.Len()
}

func prioSched() {
	prioQueue := NewPrioQueue()
	prioQueue.Enqueue("A", 1)
	prioQueue.Enqueue("B", 2)
	prioQueue.Enqueue("C", 2)
	prioQueue.Enqueue("D", 3)
	prioQueue.Dequeue()
	prioQueue.Dequeue()
	prioQueue.Enqueue("D", 3)
	prioQueue.Enqueue("B", 2)
	prioQueue.Dequeue()
	prioQueue.Dequeue()
	prioQueue.Enqueue("C", 2)
	prioQueue.Enqueue("A", 1)

	// キューが空になるまで要素を取り出して表示
	for prioQueue.Size() != 0 {
		fmt.Println(prioQueue.Dequeue()) // A, C, D, D の順に表示される
	}
}

func main() {
	prioSched()
}
