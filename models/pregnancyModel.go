// pregnancyModel.go
package models

import (
	"database/sql"
	"time"
)

var _ Model = (*Pregnancy)(nil)

type Pregnancy struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	DueDate   time.Time `json:"due_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Pregnancy) ScanRow(row *sql.Row) error {
	return row.Scan(&p.ID, &p.UserID, &p.DueDate, &p.CreatedAt, &p.UpdatedAt)

}

func (p *Pregnancy) ScanRows(rows *sql.Rows) error {
	return rows.Scan(&p.ID, &p.UserID, &p.DueDate, &p.CreatedAt, &p.UpdatedAt)

}

func (p *Pregnancy) InsertQuery() (string, []any) {

	query := `INSERT INTO pregnancies (user_id,due_date)
		  VALUES ($1, $2)
		  RETURNING id, user_id, due_date, created_at, updated_at`
	args := []any{p.UserID, p.DueDate}
	return query, args

}
