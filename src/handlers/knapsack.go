package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vanessahoamea/algorithms-api/src/solvers"
	"github.com/vanessahoamea/algorithms-api/src/utils"
)

// @Summary Solves Knapsack problem
// @Description Computes the solutions for the specified Knapsack problem instance, handling both the Binary and Fractional variants of the problem.
// @Accept json
// @Produce json
// @Param request body handlers.HandleKnapsack.requestBody true "`values` represents the list of values of each object, `weights` represents the list of weights of each object, `capacity` represents the maximum weight the knapsack can hold."
// @Success 200 {object} solvers.KnapsackResult
// @Failure 400 {object} utils.ErrorResponse
// @Router /knapsack [post]
func HandleKnapsack(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Values   []int `json:"values"`
		Weights  []int `json:"weights"`
		Capacity int   `json:"capacity"`
	}

	decoder := json.NewDecoder(r.Body)
	body := requestBody{}
	err := decoder.Decode(&body)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Could not parse request body: %v", err))
		return
	}

	solver := solvers.KnapsackSolver{}
	err = solver.Initialize(body.Values, body.Weights, body.Capacity)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("%v", err))
		return
	}

	solver.Solve()
	utils.RespondWithJSON(w, 200, solver.FormatResult())
}
