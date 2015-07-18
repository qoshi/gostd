package list

import (
	"errors"
	"math"
)

type Node struct {
	Item interface{}
	Next *Node
	Prev *Node
}

type List struct {
	head *Node
	tail *Node
	size int64
}

func _init() *List {
	l := &List{
		head: &Node{},
		tail: &Node{},
		size: 0,
	}
	l.head.Next = l.tail
	l.tail.Prev = l.head
	return l
}

func newNode(item interface{}) *Node {
	return &Node{
		Item: item,
		Next: nil,
		Prev: nil,
	}
}

func New(size int64) (*List, error) {
	if size == 0 {
		return _init(), nil
	}
	if size < 0 {
		return nil, errors.New("list size should never be negative")
	}
	head := _init()
	//to do init the length
	return head, nil
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
	l.head = &Node{}
	l.tail = &Node{}
	l.head.Next = l.tail
	l.tail.Prev = l.head
}

func (l *List) getPos(pos int64) *Node {
	p := l.head.Next
	var i int64
	for i = 0; i < pos; i++ {
		p = p.Next
	}
	return p
}

func (l *List) Insert(pos int64, val interface{}) error {
	if pos < 0 || pos > l.size {
		return errors.New("pos invalid")
	}
	node := newNode(val)
	p := l.getPos(pos)
	node.Prev = p.Prev
	node.Next = p
	p.Prev.Next = node
	p.Prev = node
	l.size++
	return nil
}

func (l *List) Erase(pos int64) error {
	if pos < 0 || pos >= l.size {
		errors.New("pos invalid")
	}
	p := l.getPos(pos)
	p.Prev.Next = p.Next
	p.Next.Prev = p.Prev
	p.Next = nil
	p.Prev = nil
	l.size--
	return nil
}

func (l *List) EraseRange(start, end int64) error {
	if start < 0 || start >= end || end > l.size {
		return errors.New("arguments invalid")
	}
	p := l.getPos(start)
	q := l.getPos(end)
	p.Prev.Next = q
	p.Prev = nil
	q.Prev.Next = nil
	q.Prev = p.Prev
	l.size -= end - start
	return nil
}

func (l *List) PushBack(item interface{}) {
	node := newNode(item)
	node.Prev = l.tail.Prev
	node.Next = l.tail
	node.Prev.Next = node
	l.tail.Prev = node
	l.size++
}

func (l *List) PopBack() error {
	if l.size == 0 {
		return errors.New("list size is zero")
	}
	p := l.tail.Prev
	l.tail.Prev = p.Prev
	p.Prev.Next = l.tail
	l.size--
	p.Next = nil
	p.Prev = nil
	return nil
}

func (l *List) PushFront(item interface{}) {
	node := newNode(item)
	node.Next = l.head.Next
	node.Prev = l.head
	l.head.Next = node
	node.Next.Prev = node
	l.size++
}

func (l *List) PopFront() error {
	if l.size == 0 {
		return errors.New("list size is zero")
	}
	p := l.head.Next
	l.head.Next = p.Next
	p.Next.Prev = l.head
	l.size--
	p.Next = nil
	p.Prev = nil
	return nil
}
