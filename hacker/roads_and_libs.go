package hacker

type Node struct {
	Value int
	Next  *Node
}

func roadsAndLibraries(n int, c_lib int, c_road int, cities [][]int) int {
	// if c_lib < c_road  =>  c_lib times number of cities
	if c_lib < c_road {
		return n * c_lib
	}

	// create array of adjacency cities
	adjList := make([]*Node, n)
	add := func(a, b int) {
		first := adjList[a]
		adjList[a] = &Node{Value: b, Next: first}
	}
	for _, pair := range cities {

		a := numToIndex(pair[0])
		b := numToIndex(pair[1])

		add(a, b)
		add(b, a)
	}

	// array of connected groups of nodes
	visited := make(map[int]struct{})
	subsets := make([][]int, 0)
	for i, n := range adjList {
		if _, ok := visited[i]; ok || n == nil {
			continue
		}
		// mark visited
		connected := make(map[int]struct{})
		reachable(adjList, connected, i)

		subset := make([]int, 0)
		for k := range connected {
			subset = append(subset, k)
			visited[k] = struct{}{}
		}
		subsets = append(subsets, subset)
	}

	// c_lib > c_road
	// for each subset
	// add 1 lib and n-1 roads
	var total int
	for _, s := range subsets {
		total += c_lib + (len(s)-1)*c_road
	}

	// add standalone nodes
	standaloneCount := n - len(visited)
	total += standaloneCount * c_lib

	return total
}

// go through list
// filter visited
// for each look up to index and repeat
func reachable(list []*Node, visited map[int]struct{}, index int) {
	_, ok := visited[index]
	if ok {
		return
	}
	n := list[index]
	if n == nil {
		return
	}

	visited[index] = struct{}{}
	reachable(list, visited, n.Value)
	visited[n.Value] = struct{}{}

	for n.Next != nil {
		n = n.Next
		if _, ok := visited[n.Value]; ok {
			continue
		}
		reachable(list, visited, n.Value)
		visited[n.Value] = struct{}{}
	}
}

func numToIndex(val int) int {
	return val - 1
}
