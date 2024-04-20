package goeasyslice_test

import (
	"runtime"
	"testing"

	eas "github.com/Matej-Chmel/go-easy-slice"
)

func check[T comparable](a, b T, t *testing.T) {
	if a != b {
		_, _, line, _ := runtime.Caller(1)
		t.Errorf("Error at line %d, \"%v\" != \"%v\"", line, a, b)
	}
}

func TestInt32(t *testing.T) {
	s := eas.NewSlice[int32](4)
	s.Set(0, 4)
	s.Set(1, 8)
	s.Set(2, 20)
	s.Set(3, 87)

	check(s.Get(0), 4, t)
	check(s.Get(1), 8, t)
	check(s.Get(2), 20, t)
	check(s.Get(3), 87, t)

	check(s.String(), "[4 8 20 87]", t)

	s.Append(17)
	s.AppendMore(78, 988, 901)

	check(s.Get(4), 17, t)
	check(s.String(), "[4 8 20 87 17 78 988 901]", t)

	s.PopVoid()
	check(s.Last(), 988, t)

	check(s.First(), 4, t)

	err := s.UpdateCapacity(2)
	check(err == nil, false, t)

	err = s.UpdateCapacity(90)
	check(err == nil, true, t)
	check(s.Cap(), 90, t)
}

func TestString(t *testing.T) {
	s := eas.NewSlice[string](0, 20)
	check(s.Cap(), 20, t)
	check(s.Len(), 0, t)
	check(s.Empty(), true, t)
	check(s.HasElements(), false, t)

	val, err := s.GetSafe(0)
	check(err == nil, false, t)
	check(val, "", t)

	s.AppendSlice([]string{"hi", "hello", "adios"})

	val, err = s.PopSafe()
	check(err == nil, true, t)
	check(val, "adios", t)

	check(s.String(), "[hi hello]", t)
}
