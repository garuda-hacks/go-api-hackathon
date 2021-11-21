package repository

import (
	"github.com/garuda-hacks/go-api-hackathon/app/model/entity"
)

type userRepository struct {
	base BaseRepository
}
type UserRepository interface {
	Find(id int) (*entity.User, error)
	FindByParam(filter map[string]interface{}) (*entity.User, error)
	Create(user entity.User) (entity.User, error)
	Update(id int, input map[string]interface{}) error
	List() ([]entity.User, error)
}

func NewUserRepository(ar BaseRepository) UserRepository {
	return &userRepository{ar}
}

func (r *userRepository) Find(id int) (*entity.User, error) {
	var user entity.User
	var baseDb = r.base

	query := baseDb.GetDB().
		Where(&entity.User{IDUser: id})

	err := query.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByParam(filter map[string]interface{}) (*entity.User, error) {
	var user entity.User
	query := r.base.GetDB()

	if filter["username"] != nil {
		query = query.Where(&entity.User{Username: filter["username"].(string)})
	}
	err := query.First(&user).Error
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (r *userRepository) Create(user entity.User) (entity.User, error) {
	err := r.base.GetDB().Create(&user).Error
	return user, err
}

func (r *userRepository) Update(id int, input map[string]interface{}) error {
	var user entity.User
	err := r.base.GetDB().Model(&user).
		Where("id_user = ?", id).
		Updates(&input).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) List() ([]entity.User, error) {
	var user []entity.User
	var baseDb = r.base

	query := baseDb.GetDB()

	err := query.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
