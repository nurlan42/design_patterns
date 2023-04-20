package usecase

import (
	"github.com/todo/internal/domain"
	"github.com/todo/internal/repository"
)

type TODOUsecase interface {
	Create(task domain.TODO) error
	GetByID(ID string) (domain.TODO, error)
	GetAll() ([]domain.TODO, error)
	UpdateByID(ID string, todo domain.TODO) error
	DeleteByID(ID string) error
}

type Usecase struct {
	TODOUsecase TODOUsecase
}

func New(r *repository.Repos) *Usecase {
	return &Usecase{
		TODOUsecase: NewTODO(r),
	}
}
