package repository

import (
	"github.com/garuda-hacks/go-api-hackathon/app/model/entity"
)

type scholarRepository struct {
	base BaseRepository
}
type ScholarRepository interface {
	Find(id int) (*entity.Scholarship, error)
	List() ([]entity.Scholarship, error)
}

func NewScholarRepository(ar BaseRepository) ScholarRepository {
	return &scholarRepository{ar}
}

func (r *scholarRepository) Find(id int) (*entity.Scholarship, error) {
	var user entity.Scholarship
	var baseDb = r.base

	query := baseDb.GetDB().
		Where(&entity.Scholarship{IDScholarship: id})

	err := query.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *scholarRepository) List() ([]entity.Scholarship, error) {
	var user []entity.Scholarship
	var baseDb = r.base

	query := baseDb.GetDB()

	err := query.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
