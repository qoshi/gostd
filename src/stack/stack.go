package stack

import "errors"

type node struct {
	Item interface{}
	Next *node
}

type Stack struct {
	front *node
	size  int64
}

func newNode(item interface{}) *node {
	return &node{
		Item: item,
		Next: nil,
	}
}

func _init() *Stack {
	return &Stack{
		front: newNode(nil),
		size:  0,
	}
}

func New() *Stack {
	return _init()
}

func (s *Stack) Top() interface{} {
	return s.front.Next.Item
}

func (s *Stack) Push(item interface{}) {
	node := newNode(item)
	node.Next = s.front.Next
	s.front.Next = node
	s.size++
}

func (s *Stack) Pop() {
	if s.front.Next == nil {
		panic(errors.New("stack is empty"))
	}
	p := s.front.Next
	s.front.Next = p.Next
	p.Next = nil
	s.size--
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Size() int64 {
	return s.size
}
