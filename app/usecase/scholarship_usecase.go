package usecase

import (
	rsp "github.com/garuda-hacks/go-api-hackathon/app/model/response"
	"github.com/garuda-hacks/go-api-hackathon/app/repository"
	"github.com/jinzhu/copier"
)

type scholarUsecase struct {
	BaseRepository     repository.BaseRepository
	ScholarRepository  repository.ScholarRepository
	UserDataRepository repository.UserDataRepository
}

type ScholarUsecase interface {
	Find(id int) (interface{}, error)
	List() (interface{}, error)
}

func NewScholarUsecase(
	br repository.BaseRepository,
	sr repository.ScholarRepository,
	udr repository.UserDataRepository,
) ScholarUsecase {
	return &scholarUsecase{
		br,
		sr,
		udr,
	}
}

func (uu *scholarUsecase) Find(id int) (interface{}, error) {
	var err error
	findscholarship, err := uu.ScholarRepository.Find(id)
	res := rsp.Scholarship{}
	copier.Copy(&res, &findscholarship)

	return res, err
}

func (uu *scholarUsecase) List() (interface{}, error) {
	// var result []interface{}
	return uu.ScholarRepository.List()
}
