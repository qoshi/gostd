package vector

import "errors"

type Vector struct {
	arr     []interface{}
	maxSize int64
	size    int64
}

func New(length int64) *Vector {
	v := Vector{
		arr:     make([]interface{}, length),
		size:    length,
		maxSize: length,
	}
	return &v
}

func (v *Vector) Get(pos int64) interface{} {
	if pos < 0 || pos >= v.size {
		panic(errors.New("pos invalid"))
	}
	return (v.arr[pos])
}

func (v *Vector) Set(pos int64, val interface{}) {
	if pos < 0 || pos >= v.size {
		panic(errors.New("pos invalid"))
	}
	v.arr[pos] = val
}

func (v *Vector) Front() interface{} {
	return v.arr[0]
}

func (v *Vector) Back() interface{} {
	if v.size > 0 {
		return v.arr[v.size-1]
	} else {
		panic(errors.New("vector is empty"))
	}
}

func (v *Vector) Empty() bool {
	return v.size == 0
}

func (v *Vector) Size() int64 {
	return v.size
}

func (v *Vector) Clear() {
	var i int64
	for i = 0; i < v.size; i++ {
		v.arr[i] = nil
	}
	v.size = 0
}

func (v *Vector) app() {
	tmp := make([]interface{}, v.maxSize)
	v.arr = append(v.arr, tmp...)
	v.maxSize = v.maxSize * 2
}

func (v *Vector) Insert(pos int64, val interface{}) {
	if pos < 0 || pos > v.size {
		panic(errors.New("pos invalid"))
	}
	if v.size+1 > v.maxSize {
		v.app()
	}
	v.size += 1
	for i := v.size; i > pos; i-- {
		v.arr[i] = v.arr[i-1]
	}
	v.arr[pos] = val
}

func (v *Vector) Erase(pos int64) {
	if v.size == 0 {
		panic(errors.New("eraser element from an empty vectory"))
	}
	if pos < 0 || pos > v.size {
		panic(errors.New("pos invalid"))
	}
	v.arr = append(v.arr[:pos], v.arr[pos+1:]...)
	v.size--
}

func (v *Vector) PushBack(val interface{}) {
	if v.size+1 > v.maxSize {
		v.app()
	}
	v.arr[v.size] = val
	v.size++
}

func (v *Vector) PopBack(val interface{}) {
	if v.size == 0 {
		panic(errors.New("pop from an empty vector"))
	}
	v.arr[v.size] = nil
	v.size--
}
