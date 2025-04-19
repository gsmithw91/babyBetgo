package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID                int            `json:"id"`
	Username          string         `json:"username"`
	Password          string         `json:"password,omitempty"`
	PasswordHash      string         `json:"-"`
	Balance           int            `json:"balance"`
	Email             sql.NullString `json:"email"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	IsActive          bool           `json:"is_active"`
	LastLogin         *time.Time     `json:"last_login,omitempty"`
	ProfilePictureURL *string        `json:"profile_picture_url,omitempty"`
	Role              string         `json:"role"`
	DisplayName       *string        `json:"display_name,omitempty"`
	Bio               *string        `json:"bio,omitempty"`
	PhoneNumber       *string        `json:"phone_number,omitempty"`
}
