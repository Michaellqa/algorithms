package roads_and_libs

import "testing"

type params struct {
	n      int
	c_lib  int
	c_road int
	cities [][]int
}

func TestRoadsAndLibraries(t *testing.T) {
	cases := []struct {
		ps  params
		res int
	}{
		{
			ps: params{
				n:      3,
				c_lib:  2,
				c_road: 1,
				cities: [][]int{
					{1, 2},
					{3, 1},
					{2, 3},
				},
			},
			res: 4,
		},
		{
			ps: params{
				n:      6,
				c_lib:  2,
				c_road: 5,
				cities: [][]int{
					{1, 3},
					{3, 4},
					{2, 4},
					{1, 2},
					{2, 3},
					{5, 6},
				},
			},
			res: 12,
		},
		{
			ps: params{
				n:      5,
				c_lib:  6,
				c_road: 1,
				cities: [][]int{
					{1, 2},
					{1, 3},
					{1, 4},
				},
			},
			res: 15,
		},
		{
			ps: params{
				n:      7,
				c_lib:  10,
				c_road: 1,
				cities: [][]int{
					{4, 1},
					{1, 3},
					{1, 2},
					{2, 3},
					{5, 6},
					{6, 7},
				},
			},
			res: 25,
		},
	}
	for _, c := range cases {
		res := roadsAndLibraries(c.ps.n, c.ps.c_lib, c.ps.c_road, c.ps.cities)
		if res != c.res {
			t.Errorf("my %d != %d expected", res, c.res)
		}
	}

}
