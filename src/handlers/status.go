package handlers

import (
	"net/http"

	"github.com/vanessahoamea/algorithms-api/src/utils"
)

func HandleStatus(w http.ResponseWriter, r *http.Request) {
	type StatusResponse struct {
		Status string `json:"status"`
	}

	utils.RespondWithJSON(w, 200, StatusResponse{
		Status: "ok",
	})
}
