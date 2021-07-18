package pqueue

import (
	"fmt"
	"testing"
)

func TestPqueueLength(t *testing.T) {
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
