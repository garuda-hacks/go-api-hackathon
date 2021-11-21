package repository

import (
	"github.com/garuda-hacks/go-api-hackathon/app/model/entity"
)

type interestRepository struct {
	base BaseRepository
}
type InterestRepository interface {
	Find(id int) ([]entity.InterestMap, error)
	FindInterest(id int) (*entity.Interest, error)
	List() ([]entity.InterestMap, error)
}

func NewInterestRepository(ar BaseRepository) InterestRepository {
	return &interestRepository{ar}
}

func (r *interestRepository) Find(id int) ([]entity.InterestMap, error) {
	var user []entity.InterestMap
	var baseDb = r.base

	query := baseDb.GetDB().
		Where(&entity.InterestMap{IDUser: id})

	err := query.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *interestRepository) FindInterest(id int) (*entity.Interest, error) {
	var user entity.Interest
	var baseDb = r.base

	query := baseDb.GetDB().
		Where(&entity.Interest{IDInterest: id})

	err := query.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *interestRepository) List() ([]entity.InterestMap, error) {
	var user []entity.InterestMap
	var baseDb = r.base

	query := baseDb.GetDB()

	err := query.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
