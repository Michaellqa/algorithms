package apple_orange

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	apples := []int{2, 3, -4}
	oranges := []int{3, -2, -4}
	a, o := countApplesAndOranges(7, 10, 4, 12, apples, oranges)
	fmt.Println(a, o)
}
