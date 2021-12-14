package tdd

import (
	"testing"
)

var nums = []int{6, 2, 3, 4, 5}

func TestInsertSort(t *testing.T) {
	result := InsertSort(nums)
	if arrayIsSorted(result) {
		t.Log("Success")
	} else {
		t.Error("slice is not sorted")
	}
}

func arrayIsSorted(numsToCheck []int) bool {
	isSorted := true
	for _, i := range numsToCheck[:len(numsToCheck)-1] {
		if i > i+1 {
			isSorted = false
			break
		}
	}

	return isSorted
}
