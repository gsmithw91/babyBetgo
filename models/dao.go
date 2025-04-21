// dao.go
package models

import "database/sql"

type Scannable interface {
	ScanRow(*sql.Row) error
	ScanRows(*sql.Rows) error
}

type Insertable interface {
	InsertQuery() (string, []any)
}

type Model interface {
	Scannable
	Insertable
}

func Insert(db *sql.DB, m Model) error {
	query, args := m.InsertQuery()
	row := db.QueryRow(query, args...)
	return m.ScanRow(row)

}

func ScanAll[T Model](rows *sql.Rows, factory func() T) ([]T, error) {

	var result []T

	for rows.Next() {
		instance := factory()
		if err := instance.ScanRows(rows); err != nil {
			return nil, err
		}
		result = append(result, instance)

	}
	return result, nil

}
