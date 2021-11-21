package repository

import (
	"github.com/garuda-hacks/go-api-hackathon/app/model/entity"
)

type messageRepository struct {
	base BaseRepository
}
type MessageRepository interface {
	List() ([]entity.Message, error)
}

func NewMessageRepository(ar BaseRepository) MessageRepository {
	return &messageRepository{ar}
}

func (r *messageRepository) List() ([]entity.Message, error) {
	var user []entity.Message
	var baseDb = r.base

	query := baseDb.GetDB()

	err := query.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
