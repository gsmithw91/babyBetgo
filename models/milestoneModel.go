// milestonesModel.go
package models

import (
	"database/sql"
	"time"
)

type Milestone struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Week        int       `json:"week"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
}

func (m *Milestone) ScanRow(row *sql.Row) error {
	return row.Scan(&m.ID, &m.UserID, &m.Week, &m.Title, &m.Description, &m.ImageURL, &m.CreatedAt)

}

func (m *Milestone) ScanRows(rows *sql.Rows) error {
	return rows.Scan(&m.ID, &m.UserID, &m.Week, &m.Title, &m.Description, &m.ImageURL, &m.CreatedAt)

}
func (m *Milestone) InsertQuery() (string, []any) {

	query := `INSERT INTO milestones (user_id,week,title,description,image_url)
		  VALUES ($1, $2)
		  RETURNING id, user_id, week, title, description, image_url, created_at`
	args := []any{m.UserID, m.Week, m.Title, m.Description, m.ImageURL}
	return query, args

}
