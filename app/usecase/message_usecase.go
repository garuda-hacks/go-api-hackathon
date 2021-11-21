package usecase

import (
	entity "github.com/garuda-hacks/go-api-hackathon/app/model/entity"
	"github.com/garuda-hacks/go-api-hackathon/app/repository"
)

type messageUsecase struct {
	BaseRepository    repository.BaseRepository
	MessageRepository repository.MessageRepository
	AccessUsecase     AccessUsecase
}

type MessageUsecase interface {
	List() ([]entity.Message, error)
}

func NewMessageUsecase(
	br repository.BaseRepository,
	mr repository.MessageRepository,
	au AccessUsecase,
) MessageUsecase {
	return &messageUsecase{
		br,
		mr,
		au,
	}
}

func (uu *messageUsecase) List() ([]entity.Message, error) {
	return uu.MessageRepository.List()
}
