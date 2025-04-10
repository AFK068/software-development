package facades

import "github.com/AFK068/bot/internal/domain"

type CategoryService interface {
	CreateCategory(category *domain.Category) error
	GetCategory(id int) (*domain.Category, error)
	DeleteCategory(id int) error
	UpdateCategory(category *domain.Category) error
}

type CategoryFacade struct {
	repo domain.Repository
}

func NewCategoryFacade(repo domain.Repository) *CategoryFacade {
	return &CategoryFacade{repo: repo}
}

func (f *CategoryFacade) CreateCategory(category *domain.Category) error {
	return f.repo.SaveCategory(category)
}

func (f *CategoryFacade) GetCategory(id int) (*domain.Category, error) {
	return f.repo.GetCategory(id)
}

func (f *CategoryFacade) DeleteCategory(id int) error {
	return f.repo.DeleteCategory(id)
}

func (f *CategoryFacade) UpdateCategory(category *domain.Category) error {
	return f.repo.UpdateCategory(category)
}
