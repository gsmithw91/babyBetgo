// babiesHandlers.go

package handlers

import (
	"babybetgo/models"
	"babybetgo/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var _ = utils.JSONResponse{}

func CreateBabyHandler(w http.ResponseWriter, r *http.Request) {

	claims, err := utils.GetClaimsFromContext(r.Context())
	if err != nil {
		utils.ErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return

	}

	pregnancyIDStr := chi.URLParam(r, "id")
	pregnancyID, err := strconv.Atoi(pregnancyIDStr)
	if err != nil || pregnancyID <= 0 {
		utils.ErrorResponse(w, "Invalid pregnancyID", http.StatusBadRequest)
		return
	}

	hasAccess, err := models.UserHasAccessToPregnancy(DB, claims.UserID, pregnancyID)
	if err != nil {
		utils.ErrorResponse(w, "Database Error", http.StatusInternalServerError)
		return
	}

	if !hasAccess {
		utils.ErrorResponse(w, "Forbidden: you don't have access to this pregnancy", http.StatusForbidden)
		return
	}

	var req struct {
		BabyName string `json:"baby_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, "Invalid Request: baby_name required", http.StatusBadRequest)
		return
	}

	baby := models.Baby{
		PregnancyID: pregnancyID,
		UserID:      claims.UserID,
		BabyName:    req.BabyName,
	}

	if err := models.Insert(DB, &baby); err != nil {
		utils.ErrorResponse(w, "Failed to create baby", http.StatusInternalServerError)
		return

	}

	utils.WriteJSON(w, http.StatusCreated, utils.JSONResponse{
		Status:  "success",
		Message: "Baby Created succesfully",
		Data:    baby,
	})
}
