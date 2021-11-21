package usecase

import (
	"errors"
	"strconv"

	req "github.com/garuda-hacks/go-api-hackathon/app/model/request"
	rsp "github.com/garuda-hacks/go-api-hackathon/app/model/response"
	"github.com/garuda-hacks/go-api-hackathon/app/repository"
	"github.com/jinzhu/copier"
)

type profilUsecase struct {
	BaseRepository     repository.BaseRepository
	UserDataRepository repository.UserDataRepository
	UgcRepository      repository.UgcRepository
	InterestRepository repository.InterestRepository
	AccessUsecase      AccessUsecase
}

type ProfilUsecase interface {
	Find(id string) (interface{}, error)
	FindParam(input req.RegProfileUser) (interface{}, error)
}

func NewProfilUsecase(
	br repository.BaseRepository,
	udr repository.UserDataRepository,
	ugc repository.UgcRepository,
	ir repository.InterestRepository,
	au AccessUsecase,
) ProfilUsecase {
	return &profilUsecase{
		br,
		udr,
		ugc,
		ir,
		au,
	}
}

func (pu *profilUsecase) Find(id string) (interface{}, error) {
	var err error
	ID, _ := strconv.Atoi(id)
	user, err := pu.UserDataRepository.Find(ID)
	ugc, _ := pu.UgcRepository.Find(ID)
	interestmap, _ := pu.InterestRepository.Find(ID)

	res := rsp.ResultProfil{}
	copier.Copy(&res, &user)

	for _, ugcdata := range ugc {
		resugc := rsp.Ugc{}
		copier.Copy(&resugc, &ugcdata)

		res.UGC = append(res.UGC, resugc)
	}

	for _, interestdata := range interestmap {
		resinterestmap := rsp.InterestMap{}
		copier.Copy(&resinterestmap, &interestdata)
		idinterest, _ := pu.InterestRepository.FindInterest(interestdata.IDInterest)
		resinterest := rsp.Interest{}
		copier.Copy(&resinterest, &idinterest)

		resinterestmap.Interest = append(resinterestmap.Interest, resinterest)
		res.Interest = append(res.Interest, resinterestmap)
	}

	respone := rsp.JsonResultProfil{}
	respone.User = res
	return respone, err
}

func (pu *profilUsecase) FindParam(input req.RegProfileUser) (interface{}, error) {
	var err error
	var res interface{}
	iduser, _ := strconv.Atoi(input.IDUser)
	check := pu.AccessUsecase.ValidToken(iduser, input.Token)
	if !check {
		err = errors.New("Unauthorized")
		return nil, err
	}
	res, err = pu.Find(input.IDUser)
	return res, err
}
