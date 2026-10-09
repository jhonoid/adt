package set

import "slices"

type Set[T comparable] struct {
	data []T
	cap  uint
}

func New[T comparable](data []T, cap uint) (s *Set[T]) {
	return
}

func (s *Set[T]) Add(elem T, elems ...T) (ok bool) {
	return
}

func (s *Set[T]) IsMember(elem T) bool {
	return slices.Contains(s.data, elem)
}

func (s *Set[T]) Get(index int) (elem T, ok bool) {
	return
}

func (s *Set[T]) IndexOf(elem T) (index int) {
	index = slices.Index(s.data, elem)
	return
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Set[T]) IsFull() bool {
	return uint(len(s.data)) == s.cap
}

func (s *Set[T]) Cardinality() uint {
	return uint(len(s.data))
}

func (s *Set[T]) Vacant() uint {
	return s.cap - uint(len(s.data))
}

func (s *Set[T]) Clear() {
}

func (s *Set[T]) Delete(elem T) {

}

func (s *Set[T]) Inter(t *Set[T]) (inter *Set[T]) {
	return
}

func (s *Set[T]) Union(t *Set[T]) (union *Set[T]) {
	return
}

func (s *Set[T]) Diff(t *Set[T]) (diff *Set[T]) {
	return
}

func (s *Set[T]) Equal(t *Set[T]) bool {
	return false
}

func (s *Set[T]) IsSubset(t *Set[T]) bool {
	return false
}

func (s *Set[T]) IsProperSubset(t *Set[T]) bool {
	return false
}

func (s *Set[T]) IsSuperset(t *Set[T]) bool {
	return false
}

func (s *Set[T]) Powerset() (ps []*Set[T]) {
	return
}

func (s *Set[T]) SymmetricDiff(t *Set[T]) (symmdiff *Set[T]) {
	return
}

func (s *Set[T]) Complement(Universal *Set[T]) (comp *Set[T]) {
	return
}

func (s *Set[T]) Cartesian(t *Set[T]) []*Set[T] {
	return nil
}
