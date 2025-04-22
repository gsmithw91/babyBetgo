// babeiesModel.go
package models

import (
	"database/sql"
	"time"
)

type Baby struct {
	ID          int       `json:"id"`
	PregnancyID int       `json:"pregnancy_id"`
	UserID      int       `json:"user_id"`
	BabyName    string    `json:"baby_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (b *Baby) ScanRow(row *sql.Row) error {
	return row.Scan(&b.ID, &b.PregnancyID, &b.UserID, &b.BabyName, &b.CreatedAt, &b.UpdatedAt)
}
func (b *Baby) ScanRows(rows *sql.Rows) error {
	return rows.Scan(&b.ID, &b.PregnancyID, &b.UserID, &b.BabyName, &b.CreatedAt, &b.UpdatedAt)
}

func (b *Baby) InsertQuery() (string, []any) {

	query := `INSERT INTO babies (pregnancy_id, user_id, baby_name)
		  VALUES ($1,$2,$3) 
		  RETURNING id, user_id, pregnancy_id, baby_name,created_at, updated_at`
	args := []any{b.PregnancyID, b.UserID, b.BabyName}
	return query, args
}
