package queue

import "list"

type Queue struct {
	l *list.List
}

func New() *Queue {
	l := list.New(0)
	return &Queue{
		l: l,
	}
}

func (q *Queue) Front() interface{} {
	return q.l.Front()
}

func (q *Queue) Back() interface{} {
	return q.l.Back()
}

func (q *Queue) Empty() bool {
	return q.l.Size() == 0
}

func (q *Queue) Size() int64 {
	return q.l.Size()
}

func (q *Queue) Push(item interface{}) {
	q.l.PushBack(item)
}

func (q *Queue) Pop() {
	q.l.PopFront()
}
