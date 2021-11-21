package usecase

import (
	rsp "github.com/garuda-hacks/go-api-hackathon/app/model/response"
	"github.com/garuda-hacks/go-api-hackathon/app/repository"
	"github.com/jinzhu/copier"
)

type internUsecase struct {
	BaseRepository       repository.BaseRepository
	InternshipRepository repository.InternshipRepository
	UserDataRepository   repository.UserDataRepository
}

type InternUsecase interface {
	Find(id int) (interface{}, error)
	List() (interface{}, error)
}

func NewInternUsecase(
	br repository.BaseRepository,
	ir repository.InternshipRepository,
	udr repository.UserDataRepository,
) InternUsecase {
	return &internUsecase{
		br,
		ir,
		udr,
	}
}

func (uu *internUsecase) Find(id int) (interface{}, error) {
	var err error
	findscholarship, err := uu.InternshipRepository.Find(id)
	res := rsp.Internship{}
	copier.Copy(&res, &findscholarship)

	return res, err
}

func (uu *internUsecase) List() (interface{}, error) {
	return uu.InternshipRepository.List()
}
