package pqueue

import (
	"fmt"
	"testing"
	"time"
)

func getMax(slice []int) int {
	max := slice[0]
	for i := 1; i < len(slice); i++ {
		if max < slice[i] {
			max = slice[i]
		}
	}
	return max
}

func TestLengthConsistency(t *testing.T) {
	scores := []int{5, 3, 7, 8, 6, 2, 9}
	manager := InitQ(len(scores))
	for _, s := range scores {
		body := fmt.Sprintf("Priority: %v", s)
		item := MakeItem(body, s)
		manager.Qpush(item)
	}
	if len(manager.Pq.List) != manager.Pq.Len() || manager.Pq.Len() != len(scores) {
		t.Fatalf(`The Queue size: "%v" is not as expected`, manager.Pq.Cursor)
	}
	length := manager.Pq.Len()
	for i := 0; i < length; i++ {
		item := manager.Qpop()
		if item == nil {
			t.Fatal("Couldn't pop item from the queue")
		}
	}
}

func TestPopNoElement(t *testing.T) {
	dummyman := InitQ(2)
	dummyman.Qpush(MakeItem("dummy", 1))
	manager := InitQ(1)
	item := manager.Qpop()
	if item != nil {
		t.Fatal("Item should be nil if popped from an empty queue")
	}

}

func TestFailToPush(t *testing.T) {
	manager := InitQ(1)
	manager.Qpush(MakeItem("Item 1", 1))
	err := manager.Qpush(MakeItem("Item 2", 5))
	if err == nil {
		t.Fatal("Qpush should return error if the queue is full")
	}
}

func TestOrder(t *testing.T) {
	scores := []int{5, 1, 17, 8, 26, 11, 8}
	manager := InitQ(len(scores))
	for _, s := range scores {
		body := fmt.Sprintf("Priority: %v", s)
		item := MakeItem(body, s)
		go manager.Qpush(item)
	}
	time.Sleep(time.Millisecond)
	length := manager.Pq.Len()
	oldPr := func(slice []int) int {
		max := slice[0]
		for i := 1; i < len(slice); i++ {
			if max < slice[i] {
				max = slice[i]
			}
		}
		return max
	}(scores)
	for i := 0; i < length; i++ {
		itm := manager.Qpop()
		ltPr := itm.Score
		if oldPr < ltPr {
			t.Fatal("Order in the queue is not consistent")
		}
		oldPr = ltPr
	}
}
