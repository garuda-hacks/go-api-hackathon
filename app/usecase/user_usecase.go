package usecase

import (
	"strconv"
	"time"

	entity "github.com/garuda-hacks/go-api-hackathon/app/model/entity"
	rsp "github.com/garuda-hacks/go-api-hackathon/app/model/response"
	"github.com/garuda-hacks/go-api-hackathon/app/repository"
	"github.com/garuda-hacks/go-api-hackathon/config"
	helper "github.com/garuda-hacks/go-api-hackathon/pkg/helper"
	"github.com/jinzhu/copier"
)

type userUsecase struct {
	BaseRepository repository.BaseRepository
	UserRepository repository.UserRepository
	AccessUsecase  AccessUsecase
}

type UserUsecase interface {
	Find(id int, token string) (interface{}, error)
	VerifyEmail(username string, password string) interface{}
	IsDuplicateEmail(username string) bool
}

func NewUserUsecase(
	br repository.BaseRepository,
	ur repository.UserRepository,
	au AccessUsecase,
) UserUsecase {
	return &userUsecase{
		br,
		ur,
		au,
	}
}

func (uu *userUsecase) Find(id int, token string) (interface{}, error) {
	var err error
	findUser, err := uu.UserRepository.Find(id)
	res := rsp.ResultUser{}
	copier.Copy(&res, &findUser)

	resp := rsp.ResponeRegister{}
	resp.User = res
	resp.Token = token
	resp.JwtToken, _ = helper.GenerateJwt(strconv.Itoa(findUser.IDUser), findUser.Username,
		config.C.Auth.ApplicationIssuer, config.C.Auth.CmsSecret, config.C.Auth.ExpiredTimeHour)
	resp.ServerTime = time.Now().Format("2006-01-02 15:04:05")

	return resp, err
}

func (uu *userUsecase) VerifyEmail(username string, password string) interface{} {
	filter := map[string]interface{}{
		"username": username,
	}
	res, err := uu.UserRepository.FindByParam(filter)
	if err != nil {
		return false
	}
	uda := entity.UserAccess{
		IDUser: res.IDUser,
	}
	token := uu.AccessUsecase.GenerateToken(uda)
	if res.Username == username && res.Password == password {
		resp, _ := uu.Find(res.IDUser, token)
		return resp
	}
	return false

}

func (uu *userUsecase) IsDuplicateEmail(username string) bool {
	filter := map[string]interface{}{
		"username": username,
	}
	res, err := uu.UserRepository.FindByParam(filter)
	if err != nil {
		return true
	}
	if res.ID == 0 {
		return true
	}
	return false
}
