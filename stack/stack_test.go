package stack

import (
	//"fmt"
	"slices"
	"testing"
)

func TestNew_CapacityInit(t *testing.T) {
	caps := []uint{3, 1, 190, 37, 0}
	for _, cp := range caps {
		s := New([]int{}, cp)
		if s.cap != cp || uint(cap(s.data)) != cp {
			t.Errorf("cap(data): %d, s.cap: %d, expected cap: %d", cap(s.data), s.cap, cp)
		}
	}
}

func TestPop_ErrEmptyStack(t *testing.T) {
	t.Parallel()

	var s Stack[int]

	_, err := s.Pop()
	if err == nil {
		t.Error("Expected non-nil error")
	}
}

func TestPop_NonNilParam(t *testing.T) {
	t.Parallel()

	wants := []int{1, 9, 32, 12, 70, 82, 77, 17}
	s := New(wants, uint(len(wants)))

	slices.Reverse(wants)
	for _, want := range wants {
		got, err := s.Pop()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if got != want {
			t.Errorf("want: %v, got: %v", want, got)
		}
	}

}

func TestUnsafePop_NonNilParam(t *testing.T) {
	t.Parallel()

	wants := []int{1, 9, 32, 12, 70, 82, 77, 17}
	s := New(wants, uint(len(wants)))

	slices.Reverse(wants)
	for _, want := range wants {
		got, err := s.UnsafePop()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if got != want {
			t.Errorf("want: %v, got: %v", want, got)
		}
	}

}

func TestPop_NilOrEmptySliceParam(t *testing.T) {
	t.Parallel()

	var s0 Stack[int]
	s1 := New[int](nil, 0)
	s2 := New([]int{}, 4)
	s3 := Stack[int]{}

	if _, err := s0.Pop(); err == nil {
		t.Errorf("VarDecl: expected non-nil error")
	}
	if _, err := s3.Pop(); err == nil {
		t.Errorf("StructDecl: expected non-nil error")
	}
	if _, err := s1.Pop(); err == nil {
		t.Errorf("NilParam: expected non-nil error")
	}
	if _, err := s2.Pop(); err == nil {
		t.Errorf("EmptySliceParam: expected non-nil error")
	}
}

func TestPush(t *testing.T) {
	t.Parallel()

	wants := []int{9, 48, 7, 43}
	s := New([]int{3, 1, 12, 90}, 8)

	//fmt.Printf("%v : cap: %v len: %v\n", s, cap(s.data), len(s.data))

	for _, want := range wants {
		err := s.Push(want)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		got, err := s.Pop()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if want != got {
			t.Errorf("want: %v, got: %v", want, got)
		}

	}
}

func TestIsEmpty(t *testing.T) {
	t.Parallel()

	s1 := New([]int{}, 0)
	s2 := New([]int{1, 2, 3}, 3)
	s3 := New[int](nil, 8)
	s4 := Stack[int]{}
	var s5 Stack[int]
	s6 := New([]int{1}, 1)

	wants := []bool{true, false, true, true, true, false}
	for i, s := range []*Stack[int]{s1, s2, s3, &s4, &s5, s6} {
		want := wants[i]
		got := s.IsEmpty()

		if want != got {
			t.Errorf("want: %t, got: %t", want, got)
		}
	}
}

func TestIsFull(t *testing.T) {
	t.Parallel()

	s1 := New([]int{}, 0)
	s2 := New([]int{1, 2, 3}, 3)
	s3 := New[int](nil, 8)
	s4 := Stack[int]{}
	var s5 Stack[int]
	s6 := New([]int{1}, 10)

	wants := []bool{true, true, false, true, true, false}
	for i, s := range []*Stack[int]{s1, s2, s3, &s4, &s5, s6} {
		want := wants[i]
		got := s.IsFull()

		if want != got {
			t.Errorf("want: %t, got: %t", want, got)
		}
	}

}

func BenchmarkPop(b *testing.B) {
	data := make([]int, 1000)
	for i := range 1000 {
		data[i] = i
	}
	s := New(data, 1000)

	for b.Loop() {
		s.Pop()
	}
}

func TestVacant(t *testing.T) {
	t.Parallel()

	s1 := New([]int{}, 0)
	s2 := New([]int{1, 2, 3}, 6)
	s3 := New[int](nil, 8)
	s4 := Stack[int]{}
	var s5 Stack[int]
	s6 := New([]int{1}, 10)

	wants := []uint{0, 3, 8, 0, 0, 9}
	for i, s := range []*Stack[int]{s1, s2, s3, &s4, &s5, s6} {
		want := wants[i]
		got := s.Vacant()

		if want != got {
			t.Errorf("want: %d, got: %d", want, got)
		}
	}

}

func TestGrow(t *testing.T) {
	t.Parallel()

	s1 := New([]int{}, 0)
	s2 := New([]int{1, 2}, 2)
	s3 := New([]int{}, 7)

	incr := []int{0, 4, 3}

	for i, s := range []*Stack[int]{s1, s2, s3} {
		want := uint(incr[i]) + s.cap
		got := s.Grow(incr[i])

		if want != got {
			t.Errorf("want: %d, got: %d", want, got)
		}
	}
}

func BenchmarkUnsafePop(b *testing.B) {
	data := make([]int, 1000)
	for i := range 1000 {
		data[i] = i
	}
	s := New(data, uint(cap(data)))

	for b.Loop() {
		s.UnsafePop()
	}
}

func BenchmarkPush(b *testing.B) {
	N := 1_000_000
	s := New([]int{}, uint(N))

	b.ResetTimer()
	for i := range N {
		err := s.Push(i)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}
