package goeasyslice

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// Generic slice with extended functionality
type EasySlice[T any] struct {
	Data []T
}

// Constructor.
// When no parameter is provided, length and capacity are both 0.
// First parameter is length, second is capacity.
func NewSlice[T any](sizes ...int) EasySlice[T] {
	if len(sizes) == 0 {
		return EasySlice[T]{make([]T, 0)}
	}

	if len(sizes) == 1 {
		return EasySlice[T]{make([]T, sizes[0])}
	}

	return EasySlice[T]{make([]T, sizes[0], sizes[1])}
}

// Adds value at the end of the slice.
func (s *EasySlice[T]) Append(value T) {
	s.Data = append(s.Data, value)
}

// Adds values at the end of the slice.
func (s *EasySlice[T]) AppendMore(values ...T) {
	s.Data = append(s.Data, values...)
}

// Adds values at the end of the slice.
func (s *EasySlice[T]) AppendSlice(values []T) {
	s.Data = append(s.Data, values...)
}

// Returns the capacity of the slice.
func (s *EasySlice[T]) Cap() int {
	return cap(s.Data)
}

// Returns bool indicating whether slice is empty.
func (s *EasySlice[T]) Empty() bool {
	return s.Len() == 0
}

// Returns copy of the first item.
func (s *EasySlice[T]) First() T {
	return s.Get(0)
}

// Returns copy of the first item or error if empty.
func (s *EasySlice[T]) FirstSafe() (T, error) {
	return s.GetSafe(0)
}

// Returns item at index i.
func (s *EasySlice[T]) Get(i int) T {
	return s.Data[i]
}

// Returns item at index i or error if i is out of bounds.
func (s *EasySlice[T]) GetSafe(i int) (T, error) {
	if i >= s.Len() {
		var none T
		return none, errors.New("out of bounds")
	}

	return s.Get(i), nil
}

// Returns bool indicating whether slice has at least one item.
func (s *EasySlice[T]) HasElements() bool {
	return !s.Empty()
}

// Returns copy of the last item.
func (s *EasySlice[T]) Last() T {
	return s.Get(s.LastIndex())
}

// Returns copy of the last item or error.
func (s *EasySlice[T]) LastSafe() (T, error) {
	return s.GetSafe(s.LastIndex())
}

// Returns last index.
func (s *EasySlice[T]) LastIndex() int {
	return s.Len() - 1
}

// Returns number of present items.
func (s *EasySlice[T]) Len() int {
	return len(s.Data)
}

// Removes and returns the last item.
func (s *EasySlice[T]) Pop() T {
	last := s.Last()
	s.PopVoid()
	return last
}

// Removes and returns the last item or returns an error.
func (s *EasySlice[T]) PopSafe() (T, error) {
	last, err := s.LastSafe()
	s.PopVoid()
	return last, err
}

// Removes last item without returning it.
func (s *EasySlice[T]) PopVoid() {
	s.Data = s.Data[:s.LastIndex()]
}

// Sets item at index i to value.
func (s *EasySlice[T]) Set(i int, value T) {
	s.Data[i] = value
}

// Sets item at index i to value or returns an error if i is out of bounds.
func (s *EasySlice[T]) SetSafe(i int, value T) error {
	if i >= s.Len() {
		return errors.New("out of bounds")
	}

	s.Data[i] = value
	return nil
}

// Returns string representation of the slice.
// Uses reflection.
func (s *EasySlice[T]) String() string {
	var builder strings.Builder
	builder.WriteRune('[')
	last := s.LastIndex()

	for i := 0; i < last; i++ {
		rVal := reflect.ValueOf(s.Get(i))
		builder.WriteString(fmt.Sprintf("%v", rVal))
		builder.WriteRune(' ')
	}

	rVal := reflect.ValueOf(s.Get(last))
	builder.WriteString(fmt.Sprintf("%v]", rVal))
	return builder.String()
}

// Attempts to update the capacity to n
// or returns an error if n less than current capacity.
func (s *EasySlice[T]) UpdateCapacity(n int) error {
	c := s.Cap()

	if n < c {
		return errors.New("cannot decrease capacity")
	}

	if n != c {
		data := make([]T, s.Len(), n)
		copy(data, s.Data)
		s.Data = data
	}

	return nil
}
