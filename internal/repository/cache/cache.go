package cache

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/todo/cfg"
	"github.com/todo/internal/domain"
	"runtime/debug"
)

type TODO struct {
	cache map[string]domain.TODO
}

func NewTODO(db *sql.DB, c *cfg.Config) *TODO {
	return &TODO{
		cache: make(map[string]domain.TODO),
	}
}

func (t *TODO) Create(td domain.TODO) error {
	td.ID = uuid.New().String()

	t.cache[td.ID] = td

	return nil
}

func (t *TODO) GetByID(ID string) (domain.TODO, error) {

	td, found := t.cache[ID]
	if !found {
		return domain.TODO{}, fmt.Errorf("GetByID: not found; %s", debug.Stack())
	}

	return td, nil
}

func (t *TODO) GetAll() ([]domain.TODO, error) {
	if len(t.cache) == 0 {
		return nil, fmt.Errorf("db empty: %s", debug.Stack())
	}
	var all []domain.TODO
	for _, val := range t.cache {
		all = append(all, val)
	}
	return all, nil
}

func (t *TODO) UpdateByID(ID string, todo domain.TODO) error {
	if _, found := t.cache[ID]; !found {
		return fmt.Errorf("UpdateByID(%s) = not found; ID = %s; %s", ID, ID, debug.Stack())
	}

	t.cache[ID] = todo

	return nil
}

func (t *TODO) DeleteByID(ID string) error {
	delete(t.cache, ID)
	return nil
}
