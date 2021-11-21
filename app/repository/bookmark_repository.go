package repository

import (
	"github.com/garuda-hacks/go-api-hackathon/app/model/entity"
)

type bookmarkRepository struct {
	base BaseRepository
}
type BookmarkRepository interface {
	List() ([]entity.Bookmark, error)
}

func NewBookmarkRepository(ar BaseRepository) BookmarkRepository {
	return &bookmarkRepository{ar}
}

func (r *bookmarkRepository) List() ([]entity.Bookmark, error) {
	var user []entity.Bookmark
	var baseDb = r.base

	query := baseDb.GetDB()

	err := query.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
