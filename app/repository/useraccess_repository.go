package repository

import (
	"time"

	"github.com/garuda-hacks/go-api-hackathon/app/model/entity"
)

type userAccessRepository struct {
	base BaseRepository
}

type UserAccessRepository interface {
	Find(id int) (entity.UserAccess, error)
	Create(ud entity.UserAccess) error
	Update(id int, token string) error
	ValidToken(id int, token string) bool
}

func NewUserAccessRepository(ar BaseRepository) UserAccessRepository {
	return &userAccessRepository{ar}
}

func (r *userAccessRepository) Find(iduser int) (entity.UserAccess, error) {
	var ua entity.UserAccess
	err := r.base.GetDB().Where("id_user = ?", iduser).
		First(&ua).Error
	if err != nil {
		return ua, err
	}
	return ua, nil
}
func (r *userAccessRepository) ValidToken(iduser int, token string) bool {
	var ua entity.UserAccess
	sql := r.base.GetDB().Where("id_user = ? and token = ?", iduser, token).
		First(&ua)
	if sql.Error != nil {
		return false
	}
	if sql.RowsAffected == 0 {
		return false
	}
	return true
}
func (r *userAccessRepository) Create(ua entity.UserAccess) error {
	ua.LastUpdate = time.Now()
	err := r.base.GetDB().Create(&ua).Error
	return err
}
func (r *userAccessRepository) Update(iduser int, token string) error {
	err := r.base.GetDB().Model(entity.UserAccess{}).
		Where("id_user = ?", iduser).
		Updates(entity.UserAccess{Token: token}).Error
	if err != nil {
		return err
	}
	return nil
}
