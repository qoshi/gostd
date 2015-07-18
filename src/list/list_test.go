package list

import (
	"fmt"
	"list"
	"testing"
)

func TestCreateList1(t *testing.T) {
	l, _ := list.New(10)
	for i := 0; i < 10; i++ {
		l.PushBack(i)
		fmt.Println(l.Size())
		fmt.Println(l.Front().(int), l.Back().(int))
	}
}

func TestCreateList2(t *testing.T) {
	l, _ := list.New(10)
	for i := 0; i < 10; i++ {
		l.PushFront(i)
		fmt.Println(l.Size())
		fmt.Println(l.Front().(int), l.Back().(int))
	}
}

func TestCreateList3(t *testing.T) {
	l, _ := list.New(10)
	for i := 0; i < 10; i++ {
		l.Insert(int64(i), i)
		fmt.Println(l.Size())
		fmt.Println(l.Front().(int), l.Back().(int))
	}
}

func TestRemove(t *testing.T) {
	l, _ := list.New(10)
	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}
	l.Erase(0)
	fmt.Println(l.Front())
	fmt.Println("sz", l.Size())
	l.Erase(1)
	fmt.Println(l.Front())
	fmt.Println("sz", l.Size())
	l.Erase(2)
	fmt.Println(l.Front())
	fmt.Println("sz", l.Size())
	for l.Size() > 0 {
		fmt.Println("sz", l.Size())
		fmt.Println(l.Front())
		l.Erase(0)
	}
}

func TestRemove2(t *testing.T) {
	l, _ := list.New(10)
	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}
	l.EraseRange(3, 8)
	fmt.Println(l.Front())
	fmt.Println("sz", l.Size())
	for l.Size() > 0 {
		fmt.Println("sz", l.Size())
		fmt.Println(l.Front())
		l.Erase(0)
	}
}

func TestPop(t *testing.T) {

}
