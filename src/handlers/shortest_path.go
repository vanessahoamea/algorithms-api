package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vanessahoamea/algorithms-api/src/solvers"
	"github.com/vanessahoamea/algorithms-api/src/utils"
)

func HandleShortestPath(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		N      int      `json:"n"`
		Edges  [][3]int `json:"edges"`
		Source int      `json:"source"`
	}

	decoder := json.NewDecoder(r.Body)
	body := requestBody{}
	err := decoder.Decode(&body)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Could not parse request body: %v", err))
		return
	}

	solver := solvers.ShortestPathSolver{}
	err = solver.Initialize(body.N, body.Edges, body.Source)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("%v", err))
		return
	}

	solver.Solve()
	utils.RespondWithJSON(w, 200, solver.FormatResult())
}
