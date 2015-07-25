package queue

import (
	"fmt"
	"queue"
	"testing"
)

func TestQ(t *testing.T) {
	q := queue.New()
	fmt.Println(q.Empty())
	fmt.Println(q.Size())
	fmt.Println("~~~~~~~~~~~")
	for i := 0; i < 10; i++ {
		q.Push(i)
		fmt.Println(q.Empty())
		fmt.Println(q.Size())
	}
	fmt.Println("~~~~~~~~~~~~~~")
	for !q.Empty() {
		fmt.Println(q.Front())
		q.Pop()
	}
}
