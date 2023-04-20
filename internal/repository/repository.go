package repository

import (
	"database/sql"

	"github.com/nurlan42/todo/cfg"
	"github.com/nurlan42/todo/internal/domain"
)

type TODORepo interface {
	Create(task domain.TODO) error
	GetByID(ID string) (domain.TODO, error)
	GetAll() ([]domain.TODO, error)
	UpdateByID(ID string, todo domain.TODO) error
	DeleteByID(ID string) error
}

type Repos struct {
	TODO TODORepo
}

func New(db *sql.DB, c *cfg.Config) *Repos {
	return &Repos{
		TODO: NewTODO(db, c),
	}
}
