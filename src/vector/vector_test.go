package vector

import (
	"fmt"
	"testing"
	"vector"
)

func TestVector(t *testing.T) {
	type Demo struct {
		IQ int64
		EQ int64
	}
	v := vector.New(10)
	fmt.Println(v.Empty())
	fmt.Println(v.Size())
	var i int64
	for i = 0; i < 10; i++ {
		fmt.Println(v.Get(i))
		tmp := Demo{
			IQ: i + 100,
			EQ: i + 200,
		}
		v.Set(i, tmp)
	}
	fmt.Println(v.Empty())
	fmt.Println(v.Size())
	fmt.Println("~~~~~~~~~~~~~~~~~~")
	fmt.Println(v.Front())
	fmt.Println(v.Back())
	fmt.Println("~~~~~~~~~~~~~~~~~~")
	v.Clear()
	fmt.Println(v.Empty())
	fmt.Println(v.Size())
	fmt.Println("~~~~~~~~~~~~~~~~~~")
	v.Insert(0, Demo{
		IQ: 100000,
		EQ: 0,
	})
	fmt.Println(v.Size())
	fmt.Println(v.Front())
	fmt.Println(v.Back())
	fmt.Println("~~~~~~~~~~~~~~~~~~")
	sb := Demo{
		IQ: 0,
		EQ: 234,
	}
	v.PushBack(sb)
	v.PushBack(sb)
	fmt.Println(v.Empty())
	fmt.Println(v.Size())
	fmt.Println(v.Front())
	fmt.Println(v.Back())
	fmt.Println("~~~~~~~~~~~~~~~~~~")
	l := v.Size()
	for i = 0; i < l; i++ {
		v.Erase(0)
		fmt.Println(i)
		fmt.Println(v.Size())
		fmt.Println(v.Empty())
		fmt.Println(v.Front())
		fmt.Println(v.Back())
	}
}
