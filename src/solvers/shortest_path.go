package solvers

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/vanessahoamea/algorithms-api/src/utils"
)

// Represents an edge pointing to a given node, having a specified weight
type edge struct {
	node   int
	weight int
}

// Represents a directed, weighted graph with no negative weights
type graph struct {
	n             int
	adjacencyList map[int][]edge
}

func (g *graph) initialize(n int, edges [][3]int) error {
	g.n = n
	g.adjacencyList = make(map[int][]edge, n)

	for i := range n {
		g.adjacencyList[i] = make([]edge, 0, n-1)
	}

	for _, group := range edges {
		if group[0] < 0 || group[0] >= n {
			return fmt.Errorf("Node %d in edge [%d, %d] (weight %d) is out of bounds. Node values belong to the interval [0, %d).", group[0], group[0], group[1], group[2], n)
		}

		if group[1] < 0 || group[1] >= n {
			return fmt.Errorf("Node %d in edge [%d, %d] (weight %d) is out of bounds. Node values belong to the interval [0, %d).", group[1], group[0], group[1], group[2], n)
		}

		if group[2] < 0 {
			return fmt.Errorf("Edge [%d, %d] has a negative weight: %d.", group[1], group[0], group[2])
		}

		edge := edge{node: group[1], weight: group[2]}
		g.adjacencyList[group[0]] = append(g.adjacencyList[group[0]], edge)
	}

	return nil
}

// Handles the problem solving logic
type ShortestPathSolver struct {
	graph     graph
	source    int
	distances []int
	previous  []int
	heap      utils.PriorityQueue[int]
}

func (s *ShortestPathSolver) Initialize(n int, edges [][3]int, source int) error {
	s.graph = graph{}
	err := s.graph.initialize(n, edges)

	s.source = source
	if source < 0 || source >= n {
		err = fmt.Errorf("Source node %d is out of bounds. Node values belong to the interval [0, %d).", source, n)
	}

	s.distances = make([]int, n)
	s.previous = make([]int, n)
	for i := range n {
		s.distances[i] = math.MaxInt
		s.previous[i] = -1
	}
	s.distances[source] = 0

	s.heap = make(utils.PriorityQueue[int], 0, n)
	heap.Init(&s.heap)

	return err
}

func (s *ShortestPathSolver) Solve() {
	s.dijkstra()
}

func (s *ShortestPathSolver) dijkstra() {
	// initialize the heap with the source node, whose minimum distance is 0
	item := utils.Item[int]{
		Value:    s.source,
		Priority: -s.distances[s.source],
	}
	heap.Push(&s.heap, &item)

	for s.heap.Len() > 0 {
		current := heap.Pop(&s.heap).(*utils.Item[int])
		node := current.Value
		distance := current.Priority * -1

		// if the node has already been processed, skip it
		if distance > s.distances[node] {
			continue
		}

		// add all adjacent neighbors to the priority queue
		for _, edge := range s.graph.adjacencyList[node] {
			if s.distances[node]+edge.weight < s.distances[edge.node] {
				s.distances[edge.node] = s.distances[node] + edge.weight
				s.previous[edge.node] = node
				newItem := utils.Item[int]{
					Value:    edge.node,
					Priority: -s.distances[edge.node],
				}
				heap.Push(&s.heap, &newItem)
			}
		}
	}
}

func (s *ShortestPathSolver) FormatResult() ShortestPathResult {
	result := ShortestPathResult{}

	result.Message = "Solution found"
	result.Solution = make([]ShortestPathResultNode, s.graph.n)
	result.FormattedOutput = ""

	for i := range s.graph.n {
		result.Solution[i] = ShortestPathResultNode{}

		result.Solution[i].Node = i
		result.Solution[i].Path = make([]int, 0, s.graph.n)

		if s.distances[i] < math.MaxInt {
			result.Solution[i].Distance = s.distances[i]

			previousNode := i
			for previousNode > -1 {
				result.Solution[i].Path = append([]int{previousNode}, result.Solution[i].Path...)
				previousNode = s.previous[previousNode]
			}

			result.FormattedOutput += fmt.Sprintf("Node %d: distance %d with path %v\n", result.Solution[i].Node, result.Solution[i].Distance, result.Solution[i].Path)
		} else {
			result.Solution[i].Distance = -1

			result.FormattedOutput += fmt.Sprintf("Node %d: Not reachable from source\n", result.Solution[i].Node)
		}
	}

	return result
}

// Represents the final solution obtained after running the algorithm
type ShortestPathResult struct {
	Message         string                   `json:"message"`
	Solution        []ShortestPathResultNode `json:"solution"`
	FormattedOutput string                   `json:"formatted_output"`
}

type ShortestPathResultNode struct {
	Node     int   `json:"node"`
	Distance int   `json:"distance"`
	Path     []int `json:"path"`
}
