// guessModel.go
package models

import (
	"database/sql"
	"time"
)

var _ Model = (*Pregnancy)(nil)

type Guess struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	PregnancyID int       `json:"pregnancy_id"`
	GenderGuess string    `json:"gender_guess"`
	CreatedAt   time.Time `json:"created_at"`
}

func (g *Guess) ScanRow(row *sql.Row) error {
	return row.Scan(&g.ID, &g.UserID, &g.PregnancyID, &g.GenderGuess, &g.CreatedAt)

}

func (g *Guess) ScanRows(rows *sql.Rows) error {

	return rows.Scan(&g.ID, &g.UserID, &g.PregnancyID, &g.GenderGuess, &g.CreatedAt)
}

func (g *Guess) InsertQuery() (string, []any) {

	query := `INSERT INTO guesses (user_id,pregnancy_id, gender_guess)
		  VALUES ($1, $2, $3)
		  RETURNING id, user_id, pregnancy_id, gender_guess, created_at`
	args := []any{g.UserID, g.PregnancyID, g.GenderGuess}
	return query, args

}
