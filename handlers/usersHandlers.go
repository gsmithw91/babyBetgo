// usersHandlers.go
package handlers

import (
	"babybetgo/models"
	"babybetgo/utils"
	"database/sql"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/go-chi/chi"
)

type PublicUserProfile struct {
	ID                int     `json:"id"`
	Username          string  `json:"username"`
	DisplayName       *string `json:"display_name,omitempty"`
	ProfilePictureURL *string `json:"profile_picture_url,omitempty"`
	Bio               *string `json:"bio,omitempty"`
	Role              string  `json:"role"`
}

func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		utils.ErrorResponse(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user := &models.User{}
	err = user.ScanRow(DB.QueryRow(`
		SELECT id, username, email, created_at, updated_at, is_active, last_login,
		       profile_picture_url, role, display_name, bio, phone_number, balance
		FROM users WHERE id = $1`, id))

	if err == sql.ErrNoRows {
		utils.ErrorResponse(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		utils.ErrorResponse(w, "Database error", http.StatusInternalServerError)
		return
	}

	resp := PublicUserProfile{
		ID:                user.ID,
		Username:          user.Username,
		DisplayName:       user.DisplayName,
		ProfilePictureURL: user.ProfilePictureURL,
		Bio:               user.Bio,
		Role:              user.Role,
	}
	jsonResp := utils.JSONResponse{
		Status:  "success",
		Message: "Retrieved public profile1",
		Data:    resp,
	}
	utils.WriteJSON(w, http.StatusOK, jsonResp)
}

func UserInfoPartialHandler(w http.ResponseWriter, r *http.Request) {
	claims, err := utils.GetClaimsFromContext(r.Context())
	if err != nil {
		utils.ErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user := &models.User{}
	err = user.ScanRow(DB.QueryRow(`
		SELECT id, username, email, balance, role, display_name 
		FROM users WHERE id = $1`, claims.UserID))

	if err == sql.ErrNoRows {
		utils.ErrorResponse(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		utils.ErrorResponse(w, "Database error", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/user_info.htmx")
	if err != nil {
		utils.ErrorResponse(w, "Template error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, user)
}

func MeHandler(w http.ResponseWriter, r *http.Request) {
	claims, err := utils.GetClaimsFromContext(r.Context())
	if err != nil {
		utils.ErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user := &models.User{}
	err = user.ScanRow(DB.QueryRow(`
		SELECT id, username, email, balance, role, display_name
		FROM users WHERE id = $1`, claims.UserID))

	if err == sql.ErrNoRows {
		utils.ErrorResponse(w, "/me route used, User not found", http.StatusNotFound)
		return
	} else if err != nil {
		utils.ErrorResponse(w, "Database error", http.StatusInternalServerError)
		return
	}

	resp := UserSummary{
		ID:          user.ID,
		Username:    user.Username,
		DisplayName: unwrapNullString(user.DisplayName),
		Role:        user.Role,
		Email:       user.Email.String,
		Balance:     user.Balance,
	}

	jsonResp := utils.JSONResponse{
		Status:  "success",
		Message: "Retrieved user details",
		Data: map[string]interface{}{
			"user":        resp,
			"server_time": time.Now(),
		},
	}
	utils.WriteJSON(w, http.StatusOK, jsonResp)

}
