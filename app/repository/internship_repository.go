package repository

import (
	"github.com/garuda-hacks/go-api-hackathon/app/model/entity"
)

type internshipRepository struct {
	base BaseRepository
}
type InternshipRepository interface {
	Find(id int) (*entity.Internship, error)
	List() ([]entity.Internship, error)
	FindByParam(filter map[string]interface{}) ([]entity.Internship, error)
}

func NewInternshipRepository(ar BaseRepository) InternshipRepository {
	return &internshipRepository{ar}
}

func (r *internshipRepository) Find(id int) (*entity.Internship, error) {
	var user entity.Internship
	var baseDb = r.base

	query := baseDb.GetDB().
		Where(&entity.Internship{IDIntern: id})

	err := query.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *internshipRepository) List() ([]entity.Internship, error) {
	var user []entity.Internship
	var baseDb = r.base

	query := baseDb.GetDB().
		Where("id_interest <= ?", "4")

	err := query.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *internshipRepository) FindByParam(filter map[string]interface{}) ([]entity.Internship, error) {
	var user []entity.Internship
	var baseDb = r.base

	query := baseDb.GetDB()

	if filter["country"] != nil {
		query = query.Where("country =?", filter["country"].(int))
	}

	err := query.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
