package usecase

import (
	"github.com/google/uuid"
	"github.com/todo/internal/domain"
	"github.com/todo/internal/repository"
)

type TODO struct {
	repos *repository.Repos
}

func NewTODO(r *repository.Repos) *TODO {
	return &TODO{
		repos: r,
	}
}

func (t *TODO) Create(td domain.TODO) error {
	td.ID = uuid.New().String()
	return t.repos.TODO.Create(td)
}

func (t *TODO) GetByID(ID string) (domain.TODO, error) {
	return t.repos.TODO.GetByID(ID)
}

func (t *TODO) GetAll() ([]domain.TODO, error) {
	return t.repos.TODO.GetAll()
}

func (t *TODO) UpdateByID(ID string, todo domain.TODO) error {
	return t.repos.TODO.UpdateByID(ID, todo)
}

func (t *TODO) DeleteByID(ID string) error {
	return t.repos.TODO.DeleteByID(ID)
}
