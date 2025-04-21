// pregnancyHandlers.go
package handlers

import (
	"babybetgo/models"
	"babybetgo/utils"
	"encoding/json"
	"log"
	"net/http"
)

var _ = utils.JSONResponse{} // compile-time usage assurance

func CreatePregnancyHandler(w http.ResponseWriter, r *http.Request) {

	claims, err := utils.GetClaimsFromContext(r.Context())
	if err != nil {
		utils.ErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req models.Pregnancy
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	req.UserID = claims.UserID

	if err := models.Insert(DB, &req); err != nil {
		log.Printf("Error Inserting pregnancy: %v", err)

		utils.ErrorResponse(w, "Error Inserting pregnancy", http.StatusBadRequest)
		return

	}

	resp := utils.JSONResponse{
		Status:  "success",
		Message: "Pregnancy created successfully",
		Data:    req,
	}

	utils.WriteJSON(w, http.StatusCreated, resp)
}
