package handlers

import (
	"babybetgo/models"
	"babybetgo/utils"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func UserProfileHandler(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	log.Println("User route hit with id: ", idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User

	err = DB.QueryRow(`SELECT id, username, email, created_at, updated_at, is_active, last_login, profile_picture_url, role, display_name, bio, phone_number, balance FROM users WHERE id=$1`, id).
		Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.IsActive, &user.LastLogin, &user.ProfilePictureURL, &user.Role, &user.DisplayName, &user.Bio, &user.PhoneNumber, &user.Balance)

	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func MeHandler(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value("user").(*utils.Claims)
	if !ok || claims == nil {
		http.Error(w, "Unauthoroized", http.StatusUnauthorized)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
