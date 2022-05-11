package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Get(id uint) (*Model, error)
	Create(model Model) (uint, error)
	Migration() error
}
type repository struct {
	db *gorm.DB
}

var _ Repository = repository{} //Compile Time Proff

func NewRepository(db *gorm.DB) Repository {
	return repository{db: db}
}

// Create implements Repository
func (repo repository) Create(model Model) (uint, error) {
	if err := repo.db.Create(&model).Error; err != nil {
		return 0, err
	}
	return model.ID, nil
}

// Get implements Repository
func (repo repository) Get(id uint) (*Model, error) {
	model := &Model{ID: id}
	if err := repo.db.First(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

// Migration implements Repository
func (repo repository) Migration() error {
	return repo.db.AutoMigrate(&Model{})
}
