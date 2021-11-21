package repository

import (
	"github.com/garuda-hacks/go-api-hackathon/app/model/entity"
)

type ugcRepository struct {
	base BaseRepository
}
type UgcRepository interface {
	Find(id int) ([]entity.Ugc, error)
	List() ([]entity.Ugc, error)
	FindByParam(filter map[string]interface{}) ([]entity.Ugc, error)
}

func NewUgcRepository(ar BaseRepository) UgcRepository {
	return &ugcRepository{ar}
}

func (r *ugcRepository) Find(id int) ([]entity.Ugc, error) {
	var user []entity.Ugc
	var baseDb = r.base

	query := baseDb.GetDB().
		Where(&entity.Ugc{IDUser: id})

	err := query.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *ugcRepository) List() ([]entity.Ugc, error) {
	var user []entity.Ugc
	var baseDb = r.base

	query := baseDb.GetDB()

	err := query.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *ugcRepository) FindByParam(filter map[string]interface{}) ([]entity.Ugc, error) {
	var user []entity.Ugc
	var baseDb = r.base

	query := baseDb.GetDB()

	if filter["interest"] != nil {
		query = query.Where("id_interest =?", filter["interest"].(int))
	}

	err := query.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
