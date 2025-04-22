// pregnancyAccessModel.go

package models

import (
	"database/sql"
	"time"
)

type PregnancyAccess struct {
	ID          int       `json:"id"`
	PregnancyID int       `json:"pregnancy_id"`
	UserID      int       `json:"user_id"`
	Role        string    `json:"role"`
	InvitedBy   *int      `json:"invited_by"`
	AccessToken *string   `json:"access_token"`
	CreatedAt   time.Time `json:"created_at"`
}

func (p *PregnancyAccess) ScanRow(row *sql.Row) error {
	return row.Scan(&p.ID, &p.PregnancyID, &p.UserID, &p.Role, &p.InvitedBy, &p.AccessToken, &p.CreatedAt)

}

func (p *PregnancyAccess) ScanRows(rows *sql.Rows) error {
	return rows.Scan(&p.ID, &p.PregnancyID, &p.UserID, &p.Role, &p.InvitedBy, &p.AccessToken, &p.CreatedAt)

}

func (p *PregnancyAccess) InsertQuery() (string, []any) {
	query := `INSERT INTO pregnancy_access 
	(pregnancy_id, user_id, role, invited_by, access_token)
	VALUES($1,$2,$3,$4,$5)
	RETURNING id, pregnancy_id, user_id, role, invited_by, access_token, created_at
	`

	args := []any{p.PregnancyID, p.UserID, p.Role, p.InvitedBy, p.AccessToken}
	return query, args
}

func UserHasAccessToPregnancy(db *sql.DB, userID int, pregnancyID int) (bool, error) {
	var exists bool
	err := db.QueryRow(`
		SELECT EXISTS (
			SELECT 1 FROM pregnancy_access
			WHERE user_id = $1 AND pregnancy_id = $2
		)
	`, userID, pregnancyID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func IsUserPregnancyOwner(db *sql.DB, userID int, pregnancyID int) (bool, error) {

	var exists bool

	err := db.QueryRow(`SELECT EXTISTS (SELECT 1 from pregnancies WHERE id = $1 AND user_id =$2)`, pregnancyID, userID).Scan(&exists)

	return exists, err

}

var _ Model = (*PregnancyAccess)(nil)
