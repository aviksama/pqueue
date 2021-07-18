package pqueue

import (
	"container/heap"
	"sync"
)

type Item struct {
	Body  string // data
	Score int    // priority
}

type PriorityQueue struct {
	List   []*Item
	Cursor int // represents length of the PriorityQueue
}

func (pq *PriorityQueue) Len() int {
	return (*pq).Cursor
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.List[i].Score > pq.List[j].Score
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.List[i], pq.List[j] = pq.List[j], pq.List[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	current_length := (*pq).Cursor
	item := x.(*Item)
	(*pq).List[current_length] = item // last existing element is at index: Cursor-1
	(*pq).Cursor = (*pq).Cursor + 1
}

func (pq *PriorityQueue) Pop() interface{} {
	current_length := (*pq).Cursor
	item := pq.List[current_length-1]
	pq.List[current_length-1] = nil
	(*pq).Cursor = (*pq).Cursor - 1
	return item
}

type QMan struct {
	Pq    *PriorityQueue
	mutex *sync.Mutex
}

func InitQ(length int) *QMan {
	qm := &QMan{}
	pq := &(PriorityQueue{
		List:   make([]*Item, length),
		Cursor: 0,
	})
	heap.Init(pq)
	qm.Pq = pq
	qm.mutex = &sync.Mutex{}
	return qm
}

func (fm *QMan) Qpush(c *Item) {
	fm.mutex.Lock()
	defer fm.mutex.Unlock()
	heap.Push(fm.Pq, c)
}

func (fm *QMan) Qpop() *Item {
	fm.mutex.Lock()
	defer fm.mutex.Unlock()
	if fm.Pq.Len() <= 0 {
		return nil
	}
	item := heap.Pop(fm.Pq).(*Item)
	return item
}

func (fm *QMan) Qremove(i int) *Item {
	fm.mutex.Lock()
	defer fm.mutex.Unlock()
	return heap.Remove(fm.Pq, i).(*Item)
}

func (fm *QMan) UpdatePriority(item *Item, newScore int) {
	fm.mutex.Lock()
	defer fm.mutex.Unlock()
	item.Score = newScore
	heap.Fix(fm.Pq, item.Score)
}
func MakeItem(body string, score int) *Item {
	return &Item{
		Body:  body,
		Score: score,
	}
}
