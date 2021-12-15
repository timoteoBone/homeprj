package main

//calculations "github.com/timoteoBone/homeprj/8-calculator"

import (
	"fmt"
)

func main() {
	nums1 := []int{5, 4, 8, 1}
	nums2 := []int{5, 4, 8, 1}
	fmt.Println("Non sorted first array", nums1)
	fmt.Println("Non sorted second array", nums2)
	insertionSort(&nums1)
	selectionSort(&nums2)
	fmt.Println("sorted first array", nums1)
	fmt.Println("sorted second array", nums2)

}

func insertionSort(nums *[]int) {
	for i := 0; i < len(*nums); i++ {
		var currentPos = i
		var temp = (*nums)[i]

		for (*nums)[currentPos-1] > temp && currentPos > 0 {
			(*nums)[currentPos] = (*nums)[currentPos-1]
			currentPos--
		}

		(*nums)[currentPos] = temp
	}
}

func selectionSort(nums *[]int) {
	for i := 0; i < len(*nums)-1; i++ {
		var min = i

		for j := i + 1; j < len(*nums); j++ {
			if (*nums)[min] > (*nums)[j] {
				min = j
			}
		}

		var temp = (*nums)[min]
		(*nums)[min] = (*nums)[i]
		(*nums)[i] = temp
	}
}
