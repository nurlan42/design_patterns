package repository

import (
	"fmt"

	"database/sql"
	"github.com/todo/cfg"
	"github.com/todo/internal/domain"
	"runtime/debug"

	_ "github.com/lib/pq"
)

type TODO struct {
	postgres *sql.DB
}

func NewTODO(db *sql.DB, c *cfg.Config) *TODO {
	return &TODO{
		postgres: db,
	}
}

func (t *TODO) Create(td domain.TODO) error {
	q := `INSERT INTO todo (id, public_id, title, description, due_date, completed) 
			VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := t.postgres.Exec(q, td.ID, td.PublicID, td.Title, td.Description, td.DueDate, td.Completed)
	if err != nil {
		return fmt.Errorf("create(): %v; %s%", err, debug.Stack())
	}

	return nil
}

func (t *TODO) GetByID(ID string) (domain.TODO, error) {
	row, err := t.postgres.Query("SELECT * from todo WHERE id = ?;", ID)
	if err != nil {
		return domain.TODO{}, err
	}

	var td domain.TODO
	if err := row.Scan(td); err != nil {
		return domain.TODO{}, fmt.Errorf("GetByID: %v; %v", err, debug.Stack())
	}

	return td, nil
}

func (t *TODO) GetAll() ([]domain.TODO, error) {
	return nil, nil
}

func (t *TODO) UpdateByID(ID string, todo domain.TODO) error {
	return nil
}

func (t *TODO) DeleteByID(ID string) error {
	return nil
}
