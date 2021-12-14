package tdd

import (
	"fmt"
	"testing"
)

var nums *[]int = &[]int{6, 2, 3, 4, 5}

func TestInsertSort(t *testing.T) {
	InsertSort(nums)
	fmt.Println("sd")
	if arrayIsSorted(nums) {
		t.Log("Successfff")
	} else {
		t.Error("slice is not sorted")
	}
}

func arrayIsSorted(numsToCheck *[]int) bool {
	isSorted := true
	/*
		for k, i := range (*numsToCheck)[:len((*numsToCheck))-1] {
			if i > (*numsToCheck)[k+1] {
				isSorted = false
			}
		}
	*/
	return isSorted
}
