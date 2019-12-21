package bfs

import "testing"

type dstForIndex struct {
	index int
	dists []int
}

func TestGraph_Distances(t *testing.T) {
	edgeWeight := 6
	cases := []struct {
		n     int
		edges [][]int
		exp   []dstForIndex
	}{
		{
			n: 6,
			edges: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{1, 5},
			},
			exp: []dstForIndex{
				{index: 0, dists: []int{0, -1, -1, -1, -1, -1}},
				{index: 1, dists: []int{-1, 0, 6, 12, 18, 6}},
				{index: 2, dists: []int{-1, 6, 0, 6, 12, 12}},
			},
		},
		{
			n: 5,
			edges: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 0},
				{1, 0},
			},
			exp: []dstForIndex{
				{index: 0, dists: []int{0, 6, 12, 12, 6}},
			},
		},
	}

	for _, c := range cases {
		graph := NewGraph(c.n, edgeWeight, c.edges)
		for _, idx := range c.exp {
			res := graph.Distances(idx.index)

			if len(res) != len(idx.dists) {
				t.Error("len")
			}
			for i := range res {
				if res[i] != idx.dists[i] {
					t.Errorf("#%d  my %d != %d expected", i, res[i], idx.dists[i])
				}
			}
		}
	}
}
