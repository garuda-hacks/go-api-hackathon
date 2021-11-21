package usecase

import (
	entity "github.com/garuda-hacks/go-api-hackathon/app/model/entity"
	"github.com/garuda-hacks/go-api-hackathon/app/repository"
	logger "github.com/garuda-hacks/go-api-hackathon/infrastructure/io"
	uuid "github.com/satori/go.uuid"
)

type accessUsecase struct {
	BaseRepository       repository.BaseRepository
	UserAccessRepository repository.UserAccessRepository
}

type AccessUsecase interface {
	GenerateToken(input entity.UserAccess) string
	ValidToken(id int, token string) bool
}

func NewAccessUsecase(
	br repository.BaseRepository,
	uar repository.UserAccessRepository,
) AccessUsecase {
	return &accessUsecase{br, uar}
}

func (uu *accessUsecase) GenerateToken(input entity.UserAccess) string {
	cl := logger.WithFields(logger.Fields{"UserInteractor": "Token"})
	var token string
	cekAcc, _ := uu.UserAccessRepository.Find(input.IDUser)
	if cekAcc.Token == "" {
		token = uuid.NewV4().String()
		input.Token = token
		err := uu.UserAccessRepository.Create(input)
		if err != nil {
			cl.Errorf("[ERROR]  %v", err.Error())
		}
	} else {
		token = uuid.NewV4().String()
		err := uu.UserAccessRepository.Update(input.IDUser, token)
		if err != nil {
			cl.Errorf("[ERROR]  %v", err.Error())
		}
	}
	return token
}
func (uu *accessUsecase) ValidToken(id int, token string) bool {
	res := uu.UserAccessRepository.ValidToken(id, token)
	return res
}
