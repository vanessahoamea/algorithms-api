package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vanessahoamea/algorithms-api/src/solvers"
	"github.com/vanessahoamea/algorithms-api/src/utils"
)

// @Summary Solves N Queens problem
// @Description Computes the solution for the specified N Queens problem instance, using the Forward Checking algorithm with MRV sorting.
// @Accept json
// @Produce json
// @Param request body handlers.HandleNQueens.requestBody true "`n` represents the number of queens (one for each row/column), `blocked` represents the blocked squares on the chessboard."
// @Success 200 {object} solvers.NQueensResult
// @Failure 400 {object} utils.ErrorResponse
// @Router /n-queens [post]
func HandleNQueens(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		N       int     `json:"n"`
		Blocked [][]int `json:"blocked"`
	}

	decoder := json.NewDecoder(r.Body)
	body := requestBody{}
	err := decoder.Decode(&body)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Could not parse request body: %v", err))
		return
	}

	solver := solvers.NQueensSolver{}
	err = solver.Initialize(body.N, body.Blocked)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("%v", err))
		return
	}

	solver.Solve()
	utils.RespondWithJSON(w, 200, solver.FormatResult())
}
