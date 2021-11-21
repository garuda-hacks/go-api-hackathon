package registry

import (
	"github.com/garuda-hacks/go-api-hackathon/app/controller"
	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewProfilController() controller.ProfilController
	NewUserController() controller.UserController
	NewScholarController() controller.ScholarController
	NewInternController() controller.InternController
	NewUgcController() controller.UgcController
	NewMessageController() controller.MessageController
	NewBookmarkController() controller.BookmarkController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}
