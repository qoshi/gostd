package stack

import (
	"fmt"
	"stack"
	"testing"
)

func TestCreateStack(t *testing.T) {
	s := stack.New()
	fmt.Println(s.Empty(), s.Size())
	for i := 0; i < 10; i++ {
		s.Push(i)
		fmt.Println(s.Size(), s.Empty(), s.Top())
	}
	for i := 0; i < 10; i++ {
		fmt.Println(s.Size(), s.Empty(), s.Top())
		s.Pop()
	}
	fmt.Println("~~~~~~")
	fmt.Println(s.Size())
	fmt.Println(s.Empty())
	fmt.Println(s.Pop())
}

func TestCreateStack1(t *testing.T) {
	s := stack.New()
	for i := 0; i < 10; i++ {
		s.Push(i)
		fmt.Println(s.Size(), s.Empty(), s.Top())
		s.Pop()
	}
	fmt.Println(s.Size())
	fmt.Println(s.Empty())
	fmt.Println(s.Pop())
}
