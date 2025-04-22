// pregnancyHandlers.go
package handlers

import (
	"babybetgo/models"
	"babybetgo/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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

	ownerAcces := &models.PregnancyAccess{
		PregnancyID: req.ID,
		UserID:      claims.UserID,
		Role:        "owner",
		InvitedBy:   &claims.UserID,
	}

	_ = models.Insert(DB, ownerAcces)

	resp := utils.JSONResponse{
		Status:  "success",
		Message: "Pregnancy created successfully",
		Data:    req,
	}

	utils.WriteJSON(w, http.StatusCreated, resp)
}

func GrantAccessHandler(w http.ResponseWriter, r *http.Request) {
	claims, err := utils.GetClaimsFromContext(r.Context())
	if err != nil {
		utils.ErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	pregnancyIDstr := chi.URLParam(r, "id")
	pregnancyID, err := strconv.Atoi(pregnancyIDstr)
	if err != nil || pregnancyID <= 0 {
		utils.ErrorResponse(w, "Invalid pregnancy ID", http.StatusBadRequest)
		return
	}

	isOwner, err := models.IsUserPregnancyOwner(DB, claims.UserID, pregnancyID)
	if err != nil {
		utils.ErrorResponse(w, "Database error", http.StatusInternalServerError)
		return
	}
	if !isOwner {
		utils.ErrorResponse(w, "Only the owner can grant access", http.StatusForbidden)
		return
	}

	var req struct {
		UserID int    `json:"user_id"`
		Role   string `json:"role"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.UserID <= 0 {
		utils.ErrorResponse(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if req.Role == "" {
		req.Role = "guesser"
	}

	exists, err := models.UserHasAccessToPregnancy(DB, req.UserID, pregnancyID)
	if err != nil {
		utils.ErrorResponse(w, "Database error", http.StatusInternalServerError)
		return
	}
	if exists {
		utils.ErrorResponse(w, "User already has access", http.StatusConflict)
		return
	}

	access := &models.PregnancyAccess{
		PregnancyID: pregnancyID,
		UserID:      req.UserID,
		Role:        req.Role,
		InvitedBy:   &claims.UserID,
	}

	if err := models.Insert(DB, access); err != nil {
		utils.ErrorResponse(w, "Failed to grant access", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.JSONResponse{
		Status:  "success",
		Message: "Access granted",
		Data:    access,
	})
}
