package main

//calculations "github.com/timoteoBone/homeprj/8-calculator"
import "fmt"

func main() {
	var nums *[]int = &[]int{5, 4, 2, 5}
	fmt.Println(nums)
	InsertSort(nums)
	fmt.Println(nums)
}

func InsertSort(nums *[]int) {
	for k, i := range (*nums)[:len((*nums))-1] {
		if i > (*nums)[k+1] {
			(*nums)[k] = (*nums)[k+1]
			(*nums)[k+1] = i

		}
		k--
	}

}
