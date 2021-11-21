package repository

import (
	"github.com/garuda-hacks/go-api-hackathon/app/model/entity"
)

type userDataRepository struct {
	base BaseRepository
}
type UserDataRepository interface {
	Find(id int) (*entity.UserData, error)
	FindByParam(filter map[string]interface{}) (*entity.UserData, error)
	Create(user entity.UserData) (entity.UserData, error)
	Update(id int, input map[string]interface{}) error
	List() ([]entity.UserData, error)
}

func NewUserDataRepository(ar BaseRepository) UserDataRepository {
	return &userDataRepository{ar}
}

func (r *userDataRepository) Find(id int) (*entity.UserData, error) {
	var user entity.UserData
	var baseDb = r.base

	query := baseDb.GetDB().
		Where(&entity.UserData{IDUser: id})

	err := query.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userDataRepository) FindByParam(filter map[string]interface{}) (*entity.UserData, error) {
	var user entity.UserData
	query := r.base.GetDB()

	if filter["first_name"] != nil {
		query = query.Where(&entity.UserData{FirstName: filter["first_name"].(string)})
	}
	if filter["id_user"] != nil {
		query = query.Where(&entity.UserData{IDUser: filter["id_user"].(int)})
	}
	err := query.First(&user).Error
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (r *userDataRepository) Create(user entity.UserData) (entity.UserData, error) {
	err := r.base.GetDB().Create(&user).Error
	return user, err
}

func (r *userDataRepository) Update(id int, input map[string]interface{}) error {
	var user entity.UserData
	err := r.base.GetDB().Model(&user).
		Where("id_user = ?", id).
		Updates(&input).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userDataRepository) List() ([]entity.UserData, error) {
	var user []entity.UserData
	var baseDb = r.base

	query := baseDb.GetDB()

	err := query.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
