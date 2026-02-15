package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vanessahoamea/algorithms-api/src/solvers"
	"github.com/vanessahoamea/algorithms-api/src/utils"
)

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
