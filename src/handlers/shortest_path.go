package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vanessahoamea/algorithms-api/src/solvers"
	"github.com/vanessahoamea/algorithms-api/src/utils"
)

// @Summary Solves Shortest Path problem
// @Description Computes the solution for the specified single-source Shortest Path problem instance, using Dijkstra's algorithm.
// @Accept json
// @Produce json
// @Param request body handlers.HandleShortestPath.requestBody true "`n` represents the number of nodes in the graph, `edges` represents the directed edges in the graph in a [start, end, weight] format, `source` represents the source node from which all paths will be calculated."
// @Success 200 {object} solvers.ShortestPathResult
// @Failure 400 {object} utils.ErrorResponse
// @Router /shortest-path [post]
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
