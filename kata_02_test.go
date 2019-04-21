package gocodecheck

import (
	"testing"
)

//TODO get unit test from website
func TestChop(t *testing.T) {
	assertEqual := func(expected int, result int) {
		if expected != result {
			t.Errorf("Error search expected %d, result %d", expected, result)
		}
	}

	assertEqual(-1, chop(3, []int{}))
	assertEqual(-1, chop(3, []int{1}))
	assertEqual(0, chop(1, []int{1}))

	assertEqual(0, chop(1, []int{1, 3, 5}))
	assertEqual(1, chop(3, []int{1, 3, 5}))
	assertEqual(2, chop(5, []int{1, 3, 5}))
	assertEqual(-1, chop(0, []int{1, 3, 5}))
	assertEqual(-1, chop(2, []int{1, 3, 5}))
	assertEqual(-1, chop(4, []int{1, 3, 5}))
	assertEqual(-1, chop(6, []int{1, 3, 5}))
}

func TestChopIterative(t *testing.T) {
	assertEqual := func(expected int, result int) {
		if expected != result {
			t.Errorf("Error search iterative expected %d, result %d", expected, result)
		}
	}

	assertEqual(-1, chopIterative(3, []int{}))
	assertEqual(-1, chopIterative(3, []int{1}))
	assertEqual(0, chopIterative(1, []int{1}))

	assertEqual(0, chopIterative(1, []int{1, 3, 5}))
	assertEqual(1, chopIterative(3, []int{1, 3, 5}))
	assertEqual(2, chopIterative(5, []int{1, 3, 5}))
	assertEqual(-1, chopIterative(0, []int{1, 3, 5}))
	assertEqual(-1, chopIterative(2, []int{1, 3, 5}))
	assertEqual(-1, chopIterative(4, []int{1, 3, 5}))
	assertEqual(-1, chopIterative(6, []int{1, 3, 5}))
}

func TestChopSlice(t *testing.T) {
	assertEqual := func(expected int, result int) {
		if expected != result {
			t.Errorf("Error search slice expected %d, result %d", expected, result)
		}
	}

	assertEqual(-1, chopSlice(3, []int{}))
	assertEqual(-1, chopSlice(3, []int{1}))
	assertEqual(0, chopSlice(1, []int{1}))

	assertEqual(0, chopSlice(1, []int{1, 3, 5}))
	assertEqual(1, chopSlice(3, []int{1, 3, 5}))
	assertEqual(2, chopSlice(5, []int{1, 3, 5}))
	assertEqual(-1, chopSlice(0, []int{1, 3, 5}))
	assertEqual(-1, chopSlice(2, []int{1, 3, 5}))
	assertEqual(-1, chopSlice(4, []int{1, 3, 5}))
	assertEqual(-1, chopSlice(6, []int{1, 3, 5}))
}
