package quicksort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{3, 14, 9, 12, 19, 6, 10, 5, 18, 11, 2, 8}
	res := Sorts(arr)
	fmt.Println(res)
}
