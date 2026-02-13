package solvers

import (
	"errors"
	"fmt"
	"slices"
	"sort"

	"github.com/vanessahoamea/algorithms-api/src/utils"
)

// Represents an item that can be added to the knapsack
type item struct {
	value  int
	weight int
}

func (i *item) compareRatio(other *item) bool {
	thisRatio := float64(i.value) / float64(i.weight)
	otherRatio := float64(other.value) / float64(other.weight)
	return otherRatio < thisRatio
}

// Represents a knapsack with a maximum capacity for storing the given items
type knapsack struct {
	capacity int
	n        int
	items    []item
}

func (k *knapsack) initialize(values []int, weights []int, capacity int) error {
	if len(values) != len(weights) {
		return fmt.Errorf("Lenght of values array (%d) does not match length of weights arrays (%d).", len(values), len(weights))
	}

	if slices.Contains(weights, 0) {
		return errors.New("Weights array can not contain null values.")
	}

	k.capacity = capacity
	k.n = len(values)
	k.items = make([]item, k.n)
	for i := range values {
		k.items[i] = item{value: values[i], weight: weights[i]}
	}

	return nil
}

// Handles the problem solving logic
type KnapsackSolver struct {
	knapsack         knapsack
	binaryItems      map[int]bool
	binaryValue      int
	binaryWeight     int
	fractionalItems  map[int]float64
	fractionalValue  float64
	fractionalWeight float64
}

func (s *KnapsackSolver) Initialize(values []int, weights []int, capacity int) error {
	s.knapsack = knapsack{}
	err := s.knapsack.initialize(values, weights, capacity)
	if err != nil {
		return err
	}

	n := s.knapsack.n

	s.binaryItems = make(map[int]bool, n)
	for i := range n {
		s.binaryItems[i] = false
	}
	s.binaryValue = 0
	s.binaryWeight = 0

	s.fractionalItems = make(map[int]float64, n)
	for i := range n {
		s.fractionalItems[i] = 0.0
	}
	s.fractionalValue = 0.0
	s.fractionalWeight = 0.0

	return nil
}

func (s *KnapsackSolver) Solve() {
	s.solveBinaryVersion()
	s.solveFractionalVersion()
}

func (s *KnapsackSolver) solveBinaryVersion() {
	n := s.knapsack.n
	capacity := s.knapsack.capacity
	items := s.knapsack.items

	table := make([][]int, n+1)
	for i := range n + 1 {
		table[i] = make([]int, capacity+1)

		for j := range capacity + 1 {
			if i == 0 || j == 0 {
				table[i][j] = 0
			} else {
				if items[i-1].weight <= j {
					includeItem := items[i-1].value + table[i-1][j-items[i-1].weight]
					excludeItem := table[i-1][j]
					table[i][j] = max(includeItem, excludeItem)
				} else {
					table[i][j] = table[i-1][j]
				}
			}
		}
	}

	currentValue := table[n][capacity]
	currentWeight := capacity
	for i := n; i > 0 && currentValue > 0; i-- {
		if currentValue != table[i-1][currentWeight] {
			s.binaryItems[i-1] = true
			s.binaryWeight += items[i-1].weight
			currentValue -= items[i-1].value
			currentWeight -= items[i-1].weight
		}
	}

	s.binaryValue = table[n][capacity]
}

func (s *KnapsackSolver) solveFractionalVersion() {
	n := s.knapsack.n
	capacity := s.knapsack.capacity

	// creating a copy of the items array
	items := make([]item, len(s.knapsack.items))
	for i := range n {
		items[i] = item{
			value:  s.knapsack.items[i].value,
			weight: s.knapsack.items[i].weight,
		}
	}

	// preserving the original order of the items before sorting, so we can build the solution
	indices := make([]int, len(items))
	for i := range n {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		return items[indices[i]].compareRatio(&items[indices[j]])
	})

	// ordering the items in decreasing order, by value-to-weight ratio
	sort.Slice(items, func(i, j int) bool {
		return items[i].compareRatio(&items[j])
	})

	for newIndex, oldIndex := range indices {
		if items[newIndex].weight <= capacity {
			s.fractionalItems[oldIndex] = 1.0
			s.fractionalValue += float64(items[newIndex].value)
			s.fractionalWeight += float64(items[newIndex].weight)
			capacity -= items[newIndex].weight
		} else {
			ratio := float64(capacity) / float64(items[newIndex].weight)
			s.fractionalItems[oldIndex] = ratio
			s.fractionalValue += float64(items[newIndex].value) * ratio
			s.fractionalWeight += float64(capacity)
			break
		}
	}
}

func (s *KnapsackSolver) FormatResult() KnapsackResult {
	result := KnapsackResult{}

	result.BinarySolution = KnapsackResultData[int]{}
	result.FractionalSolution = KnapsackResultData[float64]{}
	result.FormattedOutput = ""

	// binary version (we can only select an object in its entirety)
	result.BinarySolution.MaxValue = s.binaryValue
	result.BinarySolution.MaxWeight = s.binaryWeight
	result.BinarySolution.SelectedItems = make([]KnapsackResultItem[int], 0, s.knapsack.n)
	for i := range s.binaryItems {
		if s.binaryItems[i] {
			if len(result.BinarySolution.SelectedItems) == 0 {
				result.FormattedOutput += "Binary version:\n"
			}

			item := KnapsackResultItem[int]{
				Number: i,
				Value:  s.knapsack.items[i].value,
				Weight: s.knapsack.items[i].weight,
				Ratio:  1.0,
			}
			result.BinarySolution.SelectedItems = append(result.BinarySolution.SelectedItems, item)

			result.FormattedOutput += fmt.Sprintf("Item %d: Value = $%d, Weight = %d kg\n", item.Number, item.Value, item.Weight)
		}
	}

	hasBinarySolution := len(result.BinarySolution.SelectedItems) > 0
	if hasBinarySolution {
		result.FormattedOutput += fmt.Sprintf("-> Total value: $%d\n", result.BinarySolution.MaxValue)
		result.FormattedOutput += fmt.Sprintf("-> Total weight: %d kg (out of %d kg)\n", result.BinarySolution.MaxWeight, s.knapsack.capacity)
	}

	// fractional version (we can also select a fraction of an object)
	result.FractionalSolution.MaxValue = s.fractionalValue
	result.FractionalSolution.MaxWeight = s.fractionalWeight
	result.FractionalSolution.SelectedItems = make([]KnapsackResultItem[float64], 0, s.knapsack.n)
	for i := range s.fractionalItems {
		if s.fractionalItems[i] > 0.0 {
			if len(result.FractionalSolution.SelectedItems) == 0 {
				if hasBinarySolution {
					result.FormattedOutput += "\n"
				}
				result.FormattedOutput += "Fractional version:\n"
			}

			item := KnapsackResultItem[float64]{
				Number: i,
				Value:  float64(s.knapsack.items[i].value) * s.fractionalItems[i],
				Weight: float64(s.knapsack.items[i].weight) * s.fractionalItems[i],
				Ratio:  s.fractionalItems[i],
			}
			result.FractionalSolution.SelectedItems = append(result.FractionalSolution.SelectedItems, item)

			if utils.FloatEqual(item.Ratio, 1.0) {
				result.FormattedOutput += fmt.Sprintf("Item %d: Value = $%.2f, Weight = %.2f kg\n", item.Number, item.Value, item.Weight)
			} else {
				result.FormattedOutput += fmt.Sprintf("Item %d: Value = $%.2f, Weight = %.2f kg (%.2f%% of whole object)\n", item.Number, item.Value, item.Weight, item.Ratio*100.0)
			}
		}
	}

	hasFractionalSolution := len(result.FractionalSolution.SelectedItems) > 0
	if hasFractionalSolution {
		result.FormattedOutput += fmt.Sprintf("-> Total value: $%.2f\n", result.FractionalSolution.MaxValue)
		result.FormattedOutput += fmt.Sprintf("-> Total weight: %.2f kg (out of %d kg)\n", result.FractionalSolution.MaxWeight, s.knapsack.capacity)
	}

	if hasBinarySolution && hasFractionalSolution {
		result.Message = "Binary & Fractional solution found"
	} else if hasBinarySolution {
		result.Message = "Only Binary solution found"
	} else if hasFractionalSolution {
		result.Message = "Only Fractional solution found"
	} else {
		result.Message = "No solution"
	}

	return result
}

// Represents the final solution obtained after running the algorithm
type KnapsackResult struct {
	Message            string                      `json:"message"`
	BinarySolution     KnapsackResultData[int]     `json:"binary_solution"`
	FractionalSolution KnapsackResultData[float64] `json:"fractional_solution"`
	FormattedOutput    string                      `json:"formatted_output"`
}

type KnapsackResultData[T int | float64] struct {
	MaxValue      T                       `json:"max_value"`
	MaxWeight     T                       `json:"max_weight"`
	SelectedItems []KnapsackResultItem[T] `json:"selected_items"`
}

type KnapsackResultItem[T int | float64] struct {
	Number int     `json:"number"`
	Value  T       `json:"value"`
	Weight T       `json:"weight"`
	Ratio  float64 `json:"ratio"`
}
