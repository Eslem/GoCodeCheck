package gocodecheck

import (
	"testing"
)

//TODO get unit test from website
func TestKarateChop(t *testing.T) {
	arr1 := []int{0, 1, 2, 3, 4, 5, 6}

	sol1 := chop(2, arr1)
	sol2 := chop(100, arr1)

	if sol1 != 2 {
		t.Errorf("Error search value %d, result %d", 2, sol1)
	}

	if sol2 != -1 {
		t.Errorf("Error search value %d, result %d", 2, sol2)
	}

}
