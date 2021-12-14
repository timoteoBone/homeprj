package tdd

func InsertSort(nums *[]int) {
	for k, i := range (*nums)[:len((*nums))-1] {
		if i > (*nums)[k+1] {
			next := (*nums)[k]
			(*nums)[k] = (*nums)[k+1]
			(*nums)[k+1] = (*nums)[next]
			k--
		}
	}

}

// 5,3,6
//3
