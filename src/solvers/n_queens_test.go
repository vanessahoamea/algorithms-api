package solvers

import (
	"fmt"
	"maps"
	"testing"
)

func TestNQueens(t *testing.T) {
	type testCase struct {
		n               int
		blocked         [][]int
		expectedDomains []queenDomain
		expectedRows    []int
	}

	testCases := []testCase{
		{
			n: 4,
			blocked: [][]int{
				{0, 0}, {1, 1}, {3, 2},
			},
			expectedDomains: []queenDomain{
				{1: true, 2: true, 3: true},
				{0: true, 2: true, 3: true},
				{0: true, 1: true, 2: true},
				{0: true, 1: true, 2: true, 3: true},
			},
			expectedRows: []int{1, 3, 0, 2},
		},
		{
			n: 5,
			blocked: [][]int{
				{0, 0}, {1, 1}, {2, 2}, {3, 4}, {4, 4},
			},
			expectedDomains: []queenDomain{
				{1: true, 2: true, 3: true, 4: true},
				{0: true, 2: true, 3: true, 4: true},
				{0: true, 1: true, 3: true, 4: true},
				{0: true, 1: true, 2: true, 3: true, 4: true},
				{0: true, 1: true, 2: true},
			},
			expectedRows: []int{2, 4, 1, 3, 0},
		},
	}

	for testCount, test := range testCases {
		solver := NQueensSolver{}
		err := solver.Initialize(test.n, test.blocked)

		// validating input data
		if err != nil {
			t.Errorf("%s", err)
			continue
		}

		// validating initial board state
		for _, queen := range solver.currentChessboard.queens {
			if !maps.Equal(queen.possibleValues, test.expectedDomains[queen.col]) {
				t.Errorf("[Queen %d] The available rows do not match the ones expected.\nActual: %v\nExpected: %v", queen.col, queen.possibleValues, test.expectedDomains[queen.col])
			}
		}

		// validating solution
		solver.Solve()
		result := solver.FormatResult()

		for _, queen := range result.Solution {
			if queen.Row != test.expectedRows[queen.Col] {
				t.Errorf("[Queen %d] The assigned row does not match the one expected.\nActual: %d\nExpected: %d", queen.Col, queen.Row, test.expectedRows[queen.Col])
			}
		}

		// print solution to help with debugging
		fmt.Printf("------------------- Test %d -------------------\n", testCount+1)
		fmt.Printf("%s\n", result.FormattedOutput)
	}
}
