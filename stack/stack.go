package stack

import (
	"errors"
	"fmt"
	"slices"
	"unsafe"
)

type Stack[T any] struct {
	data []T
	cap  uint // capacity
}

var (
	ErrEmptyStack = errors.New("stack is empty")
	ErrFullStack  = errors.New("stack is full")
)

func New[T any](data []T, cap uint) *Stack[T] {
	// cap must be atleast len(data)
	if cap < uint(len(data)) {
		panic(fmt.Errorf("invalid stack capacity: cap(%v) < data len(%v)", cap, len(data)))
	}

	d := make([]T, 0, cap)

	if len(data) == 0 {
		return &Stack[T]{
			data: d,
			cap:  cap,
		}
	}

	d = append(d, data...)
	return &Stack[T]{
		data: d,
		cap:  cap,
	}
}

func (s *Stack[T]) Pop() (lastIn T, err error) {
	dataLen := len(s.data)
	if dataLen == 0 {
		err = ErrEmptyStack
		return
	}

	lastIndex := dataLen - 1
	lastIn = s.data[lastIndex]
	s.data = s.data[:lastIndex]
	return
}

func (s *Stack[T]) UnsafePop() (lastIn T, err error) {
	if len(s.data) == 0 {
		err = ErrEmptyStack
		return
	}

	lastIndex := len(s.data) - 1

	lastIn = s.data[lastIndex]

	d := unsafe.SliceData(s.data)
	x := unsafe.Slice(d, lastIndex)

	s.data = x

	return
}

func (s *Stack[T]) Push(v T) error {
	if s.cap <= uint(len(s.data)) {
		return ErrFullStack
	}

	data := append(s.data, v)

	s.data = data
	return nil
}

/*func (s *Stack[T]) UnsafePush(v T) error {
	if s.cap <= uint(len(s.data)) {
		return ErrFullStack
	}

	return nil
}*/

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) IsFull() bool {
	return !(uint(len(s.data)) < s.cap)
}

func (s *Stack[T]) Vacant() uint {
	return (s.cap - uint(len(s.data)))
}

func (s *Stack[T]) Grow(n int) uint {
	if n <= 0 {
		return s.cap
	}

	s.cap += uint(n)

	x := slices.Grow(s.data, n)
	s.data = x

	return s.cap
}
