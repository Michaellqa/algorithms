package hash

import "testing"

func TestCountTriplets(t *testing.T) {
	cases := []struct {
		arr []int
		r   int
		exp int
	}{
		{arr: []int{1, 5, 5, 25}, r: 5, exp: 2},
		{arr: []int{1, 3, 9, 9, 27, 81}, r: 3, exp: 6},
		{arr: []int{1, 5, 5, 25, 125}, r: 5, exp: 4},
		{arr: []int{1, 5, 5, 25, 1, 1, 5, 25, 125}, r: 5, exp: 12},
	}

	for _, c := range cases {
		res := countTriplets(c.arr, c.r)
		if res != c.exp {
			t.Errorf("%d != %d", res, c.exp)
		}
	}
}
