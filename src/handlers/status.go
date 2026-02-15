package handlers

import (
	"net/http"

	"github.com/vanessahoamea/algorithms-api/src/utils"
)

// @Summary Returns the health status of the server
// @Description Used for testing purposes, when the user wants to check whether the server is currently running.
// @Produce json
// @Success 200 {object} handlers.HandleStatus.StatusResponse
// @Router /status [get]
func HandleStatus(w http.ResponseWriter, r *http.Request) {
	type StatusResponse struct {
		Status string `json:"status"`
	}

	utils.RespondWithJSON(w, 200, StatusResponse{
		Status: "ok",
	})
}
