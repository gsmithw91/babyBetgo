package handlers

import (
	"babybetgo/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username          string  `json:"username"`
	Password          string  `json:"password"`
	Email             string  `json:"email"`
	PhoneNumber       *string `json:"phonenumber"`
	Bio               *string `json:"bio"`
	DisplayName       *string `json:"display_name"`
	ProfilePictureURL *string `json:"profile_picture_url"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	db := DB

	// Check for existing username or email
	var exists bool
	err = db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM users WHERE username = $1 OR email = $2)",
		req.Username, req.Email,
	).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "Username or email already taken", http.StatusConflict)
		return
	}

	// Convert optional fields to sql.NullString
	phone := sql.NullString{Valid: false}
	if req.PhoneNumber != nil && *req.PhoneNumber != "" {
		phone = sql.NullString{String: *req.PhoneNumber, Valid: true}
	}

	bio := sql.NullString{Valid: false}
	if req.Bio != nil && *req.Bio != "" {
		bio = sql.NullString{String: *req.Bio, Valid: true}
	}

	displayName := sql.NullString{Valid: false}
	if req.DisplayName != nil && *req.DisplayName != "" {
		displayName = sql.NullString{String: *req.DisplayName, Valid: true}
	}

	profilePic := sql.NullString{Valid: false}
	if req.ProfilePictureURL != nil && *req.ProfilePictureURL != "" {
		profilePic = sql.NullString{String: *req.ProfilePictureURL, Valid: true}
	}

	_, err = db.Exec(`
		INSERT INTO users (
			username, password_hash, balance, email, role, is_active,
			created_at, updated_at, phone_number, bio, display_name, profile_picture_url
		)
		VALUES ($1, $2, 0, $3, 'user', TRUE, now(), now(), $4, $5, $6, $7)
	`,
		req.Username,
		string(hash),
		req.Email,
		phone,
		bio,
		displayName,
		profilePic,
	)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			http.Error(w, "Username or email already taken", http.StatusConflict)
			return
		}
		http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	db := DB
	var user models.User
	err := db.QueryRow(`SELECT id, username, password_hash, balance, email, created_at, updated_at, is_active, last_login, profile_picture_url, role, display_name, bio, phone_number FROM users WHERE username = $1`, req.Username).Scan(
		&user.ID, &user.Username, &user.PasswordHash, &user.Balance, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.IsActive, &user.LastLogin, &user.ProfilePictureURL, &user.Role, &user.DisplayName, &user.Bio, &user.PhoneNumber,
	)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	} else if err != nil {
		fmt.Println("Database error during login:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
