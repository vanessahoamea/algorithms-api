package solvers

import (
	"fmt"
	"testing"
)

func TestShortestPath(t *testing.T) {
	type testCase struct {
		n                 int
		edges             [][3]int
		source            int
		expectedDistances []int
		expectedPaths     [][]int
	}

	testCases := []testCase{
		{
			n: 6,
			edges: [][3]int{
				{0, 1, 2}, {0, 2, 4},
				{1, 2, 1}, {1, 3, 7},
				{2, 4, 3},
				{3, 5, 1},
				{4, 3, 2}, {4, 5, 5},
			},
			source:            0,
			expectedDistances: []int{0, 2, 3, 8, 6, 9},
			expectedPaths: [][]int{
				{0},
				{0, 1},
				{0, 1, 2},
				{0, 1, 2, 4, 3},
				{0, 1, 2, 4},
				{0, 1, 2, 4, 3, 5},
			},
		},
		{
			n: 6,
			edges: [][3]int{
				{0, 1, 50}, {0, 2, 45}, {0, 3, 10},
				{1, 2, 10}, {1, 3, 15},
				{2, 4, 30},
				{3, 0, 10}, {3, 4, 15},
				{4, 1, 20}, {4, 2, 35},
				{5, 4, 3},
			},
			source:            0,
			expectedDistances: []int{0, 45, 45, 10, 25, -1},
			expectedPaths: [][]int{
				{0},
				{0, 3, 4, 1},
				{0, 2},
				{0, 3},
				{0, 3, 4},
				{},
			},
		},
	}

	for testCount, test := range testCases {
		solver := ShortestPathSolver{}
		err := solver.Initialize(test.n, test.edges, test.source)

		// validating input data
		if err != nil {
			t.Errorf("%s", err)
			continue
		}

		// validating initial graph state
		expectedEdges := make(map[int][]edge, test.n)
		for _, group := range test.edges {
			if expectedEdges[group[0]] == nil {
				expectedEdges[group[0]] = make([]edge, 0, test.n-1)
			}
			expectedEdges[group[0]] = append(expectedEdges[group[0]], edge{node: group[1], weight: group[2]})
		}

		for i := range solver.graph.n {
			if !validateArray(solver.graph.adjacencyList[i], expectedEdges[i]) {
				t.Errorf("[Node %d] The actual adjacency list does not match the one expected.\nActual %v\nExpected: %v", i, solver.graph.adjacencyList[i], expectedEdges[i])
			}
		}

		// validating solution
		solver.Solve()
		result := solver.FormatResult()

		for i, node := range result.Solution {
			if node.Distance != test.expectedDistances[i] {
				t.Errorf("[Node %d] The actual distance does not match the one expected.\nActual %d\nExpected: %d", node.Node, node.Distance, test.expectedDistances[i])
			}
			if !validateArray(node.Path, test.expectedPaths[i]) {
				t.Errorf("[Node %d] The actual path does not match the one expected.\nActual %v\nExpected: %v", node.Node, node.Path, test.expectedPaths[i])
			}
		}

		// print solution to help with debugging
		fmt.Printf("------------------- Test %d -------------------\n", testCount+1)
		fmt.Printf("%s\n", result.FormattedOutput)
	}
}

func validateArray[T comparable](actualArray, expectedArray []T) bool {
	if len(actualArray) != len(expectedArray) {
		return false
	}

	for i, value := range expectedArray {
		if actualArray[i] != value {
			return false
		}
	}

	return true
}
