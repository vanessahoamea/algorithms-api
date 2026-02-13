package solvers

import (
	"fmt"
	"testing"

	"github.com/vanessahoamea/algorithms-api/src/utils"
)

func TestKnapsack(t *testing.T) {
	type testCase struct {
		values                   []int
		weights                  []int
		capacity                 int
		expectedBinaryItems      []int
		expectedBinaryValue      int
		expectedBinaryWeight     int
		expectedFractionalItems  []int
		expectedFractionalValue  float64
		expectedFractionalWeight float64
	}

	testCases := []testCase{
		{
			values:                   []int{19, 4, 1, 16, 16},
			weights:                  []int{32, 37, 24, 49, 27},
			capacity:                 87,
			expectedBinaryItems:      []int{0, 2, 4},
			expectedBinaryValue:      36,
			expectedBinaryWeight:     83,
			expectedFractionalItems:  []int{0, 3, 4},
			expectedFractionalValue:  44.14,
			expectedFractionalWeight: 87.00,
		},
		{
			values:                   []int{17, 9, 14, 15, 3},
			weights:                  []int{41, 36, 24, 15, 35},
			capacity:                 31,
			expectedBinaryItems:      []int{3},
			expectedBinaryValue:      15,
			expectedBinaryWeight:     15,
			expectedFractionalItems:  []int{2, 3},
			expectedFractionalValue:  24.33,
			expectedFractionalWeight: 31.00,
		},
	}

	for testCount, test := range testCases {
		solver := KnapsackSolver{}
		err := solver.Initialize(test.values, test.weights, test.capacity)

		// validating input data
		if err != nil {
			t.Errorf("%s", err)
			continue
		}

		// validating solutions
		solver.Solve()
		result := solver.FormatResult()

		if !validateSelectedItems(result.BinarySolution.SelectedItems, test.expectedBinaryItems) {
			itemNumbers := make([]int, len(result.BinarySolution.SelectedItems))
			for i, item := range result.BinarySolution.SelectedItems {
				itemNumbers[i] = item.Number
			}
			t.Errorf("[Binary solution] The selected items do not match the ones expected.\nActual: %d\nExpected: %d", itemNumbers, test.expectedBinaryItems)
		}
		if result.BinarySolution.MaxValue != test.expectedBinaryValue {
			t.Errorf("[Binary solution] The total value does not match the one expected.\nActual: %d\nExpected: %d", result.BinarySolution.MaxValue, test.expectedBinaryValue)
		}
		if result.BinarySolution.MaxWeight != test.expectedBinaryWeight {
			t.Errorf("[Binary solution] The total weight does not match the one expected.\nActual: %d\nExpected: %d", result.BinarySolution.MaxWeight, test.expectedBinaryWeight)
		}

		if !validateSelectedItems(result.FractionalSolution.SelectedItems, test.expectedFractionalItems) {
			itemNumbers := make([]int, len(result.FractionalSolution.SelectedItems))
			for i, item := range result.FractionalSolution.SelectedItems {
				itemNumbers[i] = item.Number
			}
			t.Errorf("[Fractional solution] The selected items do not match the ones expected.\nActual: %d\nExpected: %d", itemNumbers, test.expectedFractionalItems)
		}
		if !utils.FloatEqual(result.FractionalSolution.MaxValue, test.expectedFractionalValue) {
			t.Errorf("[Fractional solution] The total value does not match the one expected.\nActual: %.2f\nExpected: %.2f", result.FractionalSolution.MaxValue, test.expectedFractionalValue)
		}
		if !utils.FloatEqual(result.FractionalSolution.MaxWeight, test.expectedFractionalWeight) {
			t.Errorf("[Fractional solution] The total weight does not match the one expected.\nActual: %.2f\nExpected: %.2f", result.FractionalSolution.MaxWeight, test.expectedFractionalWeight)
		}

		// print solution to help with debugging
		fmt.Printf("------------------- Test %d -------------------\n", testCount+1)
		fmt.Printf("%s\n", result.FormattedOutput)
	}
}

func validateSelectedItems[T int | float64](actualItems []KnapsackResultItem[T], expectedItems []int) bool {
	itemsMap := make(map[int]bool, len(expectedItems))
	for _, item := range expectedItems {
		itemsMap[item] = true
	}

	validatedCount := 0
	for _, item := range actualItems {
		if itemsMap[item.Number] {
			validatedCount++
		} else {
			return false
		}
	}

	return validatedCount == len(expectedItems)
}
