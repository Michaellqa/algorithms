package bfs

type Node struct {
	Value int
	Next  *Node
}

type Graph struct {
	List       []*Node
	EdgeWeight int
}

func NewGraph(numOfNodes, edgeWeight int, edges [][]int) *Graph {
	g := &Graph{EdgeWeight: edgeWeight}
	g.List = make([]*Node, numOfNodes)
	add := func(a, b int) {
		first := g.List[a]
		g.List[a] = &Node{Value: b, Next: first}
	}
	addBidirect := func(a, b int) {
		add(a, b)
		add(b, a)
	}
	for _, e := range edges {
		addBidirect(e[0], e[1])
	}
	return g
}

func (g *Graph) Distances(startIndex int) []int {
	dists := make([]int, len(g.List))
	for i := 0; i < len(g.List); i++ {
		dists[i] = -1
	}

	// can use slice for visited, map is for convenience
	visited := make(map[int]struct{})
	frontier := map[int]struct{}{
		startIndex: {},
	}
	visited[startIndex] = struct{}{}

	// while have new nodes
	// for each reachable
	// - hasn't been visited
	// - do relaxation
	// - create new set of nodes, that hasn't been visited

	level := 0
	for len(frontier) > 0 {
		nextFrontier := make(map[int]struct{})
		for index := range frontier {
			// save distances
			dists[index] = level * g.EdgeWeight

			// find next frontier
			first := g.List[index]
			node := &Node{Next: first}

			for node.Next != nil {
				node = node.Next
				if _, ok := visited[node.Value]; ok {
					continue
				}
				nextFrontier[node.Value] = struct{}{}
				visited[node.Value] = struct{}{}
			}
		}
		frontier = nextFrontier
		level++
	}

	return dists
}
