package buffalo

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Buffalo, error)
	FindById(ID int) (Buffalo, error)
	Create(buffalo Buffalo) (Buffalo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Buffalo, error) {
	var buffalo []Buffalo
	err := r.db.Find(&buffalo).Error
	return buffalo, err
}

func (r *repository) FindById(ID int) (Buffalo, error) {
	var buffalo Buffalo
	err := r.db.Find(&buffalo).Error
	return buffalo, err
}

func (r *repository) Create(buffalo Buffalo) (Buffalo, error) {
	err := r.db.Create(&buffalo).Error
	return buffalo, err
}
