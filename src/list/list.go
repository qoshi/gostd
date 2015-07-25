package list

import (
	"errors"
	"math"
)

type node struct {
	Item interface{}
	Next *node
	Prev *node
}

type List struct {
	head *node
	tail *node
	size int64
}

func newNode(item interface{}) *node {
	return &node{
		Item: item,
		Next: nil,
		Prev: nil,
	}
}

func _init() *List {
	l := &List{
		head: newNode(nil),
		tail: newNode(nil),
		size: 0,
	}
	l.head.Next = l.tail
	l.tail.Prev = l.head
	return l
}

func New(size int64) *List {
	if size == 0 {
		return _init()
	}
	if size < 0 {
		panic(errors.New("list size should never be negative"))
	}
	head := _init()
	//to do init the length
	return head
}

func (l *List) Front() interface{} {
	if l.head.Next == nil {
		return nil
	}
	return l.head.Next.Item
}

func (l *List) Back() interface{} {
	if l.tail.Prev == nil {
		return nil
	}
	return l.tail.Prev.Item
}

func (l *List) Empty() bool {
	return l.size == 0
}

func (l *List) Size() int64 {
	return l.size
}

func (l *List) MaxSize() int64 {
	return math.MaxInt64
}

func (l *List) Clear() {
	l.size = 0
	l.head = &node{}
	l.tail = &node{}
	l.head.Next = l.tail
	l.tail.Prev = l.head
}

func (l *List) getPos(pos int64) *node {
	p := l.head.Next
	var i int64
	for i = 0; i < pos; i++ {
		p = p.Next
	}
	return p
}

func (l *List) Insert(pos int64, val interface{}) {
	if pos < 0 || pos > l.size {
		panic(errors.New("pos invalid"))
	}
	node := newNode(val)
	p := l.getPos(pos)
	node.Prev = p.Prev
	node.Next = p
	p.Prev.Next = node
	p.Prev = node
	l.size++
}

func (l *List) Erase(pos int64) {
	if pos < 0 || pos >= l.size {
		panic(errors.New("pos invalid"))
	}
	p := l.getPos(pos)
	p.Prev.Next = p.Next
	p.Next.Prev = p.Prev
	p.Next = nil
	p.Prev = nil
	l.size--
}

func (l *List) EraseRange(start, end int64) {
	if start < 0 || start >= end || end > l.size {
		panic(errors.New("arguments invalid"))
	}
	p := l.getPos(start)
	q := l.getPos(end)
	p.Prev.Next = q
	p.Prev = nil
	q.Prev.Next = nil
	q.Prev = p.Prev
	l.size -= end - start
}

func (l *List) PushBack(item interface{}) {
	node := newNode(item)
	node.Prev = l.tail.Prev
	node.Next = l.tail
	node.Prev.Next = node
	l.tail.Prev = node
	l.size++
}

func (l *List) PopBack() {
	if l.size == 0 {
		panic(errors.New("list size is zero"))
	}
	p := l.tail.Prev
	l.tail.Prev = p.Prev
	p.Prev.Next = l.tail
	l.size--
	p.Next = nil
	p.Prev = nil
}

func (l *List) PushFront(item interface{}) {
	node := newNode(item)
	node.Next = l.head.Next
	node.Prev = l.head
	l.head.Next = node
	node.Next.Prev = node
	l.size++
}

func (l *List) PopFront() {
	if l.size == 0 {
		panic(errors.New("list size is zero"))
	}
	p := l.head.Next
	l.head.Next = p.Next
	p.Next.Prev = l.head
	l.size--
	p.Next = nil
	p.Prev = nil
}
