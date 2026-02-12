package solvers

import (
	"fmt"
	"maps"
	"sort"
)

// Represents a piece to be placed on the chessboard
type queen struct {
	row            int
	col            int
	possibleValues queenDomain
}

type queenDomain = map[int]bool

func (q *queen) initialize(col int, possibleValues queenDomain) {
	q.row = -1
	q.col = col
	q.possibleValues = possibleValues
}

func (q *queen) cloneDeep() queen {
	copy := queen{
		row: q.row,
		col: q.col,
	}

	if q.possibleValues != nil {
		copy.possibleValues = make(queenDomain, len(q.possibleValues))
		maps.Copy(copy.possibleValues, q.possibleValues)
	}

	return copy
}

func (q *queen) toString() string {
	return fmt.Sprintf("(%d, %d)", q.row, q.col)
}

// Represents a chessboard instance
type chessboard struct {
	n      int
	queens []queen
}

func (c *chessboard) initialize(n int, blocked [][]int) error {
	c.n = n
	c.queens = make([]queen, n)

	for i := range n {
		possibleValues := make(queenDomain)
		for j := range n {
			possibleValues[j] = true
		}

		for _, pair := range blocked {
			if pair[0] < 0 || pair[0] >= n || pair[1] < 0 || pair[1] >= n {
				return fmt.Errorf("Blocked pair is out of bounds: [%d, %d]. Row and column values belong to the interval [0, %d).", pair[0], pair[1], n)
			}
			if pair[1] == i {
				delete(possibleValues, pair[0])
			}
		}

		queen := queen{}
		queen.initialize(i, possibleValues)
		c.queens[i] = queen
	}

	return nil
}

func (c *chessboard) cloneDeep() chessboard {
	copy := chessboard{
		n: c.n,
	}

	if c.queens != nil {
		copy.queens = make([]queen, len(c.queens))
		for i := range c.queens {
			copy.queens[i] = c.queens[i].cloneDeep()
		}
	}

	return copy
}

// Handles the problem solving logic
type NQueensSolver struct {
	currentChessboard chessboard
	queensOrder       map[int]int
	queenDomains      map[int]queenDomain
	iterations        int
	solvable          bool
}

func (s *NQueensSolver) Initialize(n int, blocked [][]int) error {
	chessboard := chessboard{}
	err := chessboard.initialize(n, blocked)

	s.currentChessboard = chessboard
	s.queensOrder = make(map[int]int)
	s.queenDomains = make(map[int]queenDomain)
	s.iterations = 0
	s.solvable = true

	return err
}

func (s *NQueensSolver) Solve() {
	s.forwardChecking()
}

func (s *NQueensSolver) forwardChecking() {
	n := s.currentChessboard.n
	beforeInitialization := make(map[int]chessboard)
	verifiedQueens := 0

	for {
		if verifiedQueens < 0 || verifiedQueens >= n {
			break
		}

		// selecting the queen with the fewest possible rows (MRV sorting)
		sortedCopy := s.currentChessboard.cloneDeep()
		sort.Slice(sortedCopy.queens, func(i, j int) bool {
			return len(sortedCopy.queens[i].possibleValues) < len(sortedCopy.queens[j].possibleValues)
		})

		var col int
		for i := range n {
			if sortedCopy.queens[i].row == -1 {
				col = sortedCopy.queens[i].col
				break
			}
		}

		row := s.selectValue(&s.currentChessboard, col)
		nextChessboard := s.placeQueen(&s.currentChessboard, col, row)

		if row == -1 {
			verifiedQueens--
			if verifiedQueens >= 0 {
				board, exists := beforeInitialization[verifiedQueens]
				if exists {
					s.currentChessboard = board.cloneDeep()
				}
			}
		} else {
			// saving the order in which the queens were processed + their domains at the time of selection
			s.queensOrder[verifiedQueens] = col
			s.queenDomains[verifiedQueens] = make(queenDomain)
			maps.Copy(s.queenDomains[verifiedQueens], s.currentChessboard.queens[col].possibleValues)

			// removing the selected value from the queen's domain, to be able to backtrack if needed
			delete(s.currentChessboard.queens[col].possibleValues, row)
			beforeInitialization[verifiedQueens] = s.currentChessboard.cloneDeep()

			verifiedQueens++
			s.currentChessboard = nextChessboard
		}

		s.iterations++
	}

	if verifiedQueens == -1 {
		s.solvable = false
	}
}

func (s *NQueensSolver) selectValue(chessboard *chessboard, col int) int {
	chessboardCopy := chessboard.cloneDeep()
	queen := chessboardCopy.queens[col]

	for len(queen.possibleValues) > 0 {
		var value int
		for key := range queen.possibleValues {
			value = key
			break
		}
		emptyDomain := false

		nextChessboard := s.placeQueen(&chessboardCopy, col, value)
		for _, newQueen := range nextChessboard.queens {
			if len(newQueen.possibleValues) == 0 && newQueen.row == -1 {
				emptyDomain = true
				break
			}
		}

		if !emptyDomain {
			return value
		}

		delete(queen.possibleValues, value)
	}

	return -1
}

func (s *NQueensSolver) placeQueen(chessboard *chessboard, col int, row int) chessboard {
	newChessboard := chessboard.cloneDeep()

	s.validateQueen(&newChessboard, col, row)

	return newChessboard
}

func (s *NQueensSolver) validateQueen(chessboard *chessboard, col int, row int) bool {
	n := chessboard.n
	queens := chessboard.queens

	if !queens[col].possibleValues[row] {
		return false
	}

	queens[col].row = row

	for i := range n {
		if queens[i].possibleValues[row] && queens[i].row == -1 {
			delete(queens[i].possibleValues, row)
		}

		// main diagonal
		if row+i < n && col+i < n && queens[col+i].possibleValues[row+i] {
			if queens[col+i].row == -1 {
				delete(queens[col+i].possibleValues, row+i)
			}
		}
		if row-i >= 0 && col-i >= 0 && queens[col-i].possibleValues[row-i] {
			if queens[col-i].row == -1 {
				delete(queens[col-i].possibleValues, row-i)
			}
		}

		// anti-diagonal
		if row+i < n && col-i >= 0 && queens[col-i].possibleValues[row+i] {
			if queens[col-i].row == -1 {
				delete(queens[col-i].possibleValues, row+i)
			}
		}
		if row-i >= 0 && col+i < n && queens[col+i].possibleValues[row-i] {
			if queens[col+i].row == -1 {
				delete(queens[col+i].possibleValues, row-i)
			}
		}
	}

	return true
}

func (s *NQueensSolver) FormatResult() NQueensResult {
	result := NQueensResult{}

	result.Iterations = s.iterations
	result.Solution = make([]NQueensResultQueen, s.currentChessboard.n)
	result.FormattedOutput = ""

	if s.solvable {
		result.Message = "Solution found"

		for index := range s.currentChessboard.n {
			queenColumn := s.queensOrder[index]
			queenDomain := make([]int, 0, len(s.queenDomains[index]))
			for key, value := range s.queenDomains[index] {
				if value {
					queenDomain = append(queenDomain, key)
				}
			}

			result.Solution[index] = NQueensResultQueen{
				Col:    queenColumn,
				Row:    s.currentChessboard.queens[queenColumn].row,
				Domain: queenDomain,
			}

			result.FormattedOutput += fmt.Sprintf("Queen %d: %v -> %s\n", queenColumn, queenDomain, s.currentChessboard.queens[queenColumn].toString())
		}
	} else {
		result.Message = "No solution"
	}

	return result
}

// Represents the final solution obtained after running the algorithm
type NQueensResult struct {
	Message         string               `json:"message"`
	Iterations      int                  `json:"iterations"`
	Solution        []NQueensResultQueen `json:"solution"`
	FormattedOutput string               `json:"formatted_output"`
}

type NQueensResultQueen struct {
	Col    int   `json:"col"`
	Row    int   `json:"row"`
	Domain []int `json:"domain"`
}
