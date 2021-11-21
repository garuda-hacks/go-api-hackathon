package usecase

import (
	rsp "github.com/garuda-hacks/go-api-hackathon/app/model/response"
	"github.com/garuda-hacks/go-api-hackathon/app/repository"
	logger "github.com/garuda-hacks/go-api-hackathon/infrastructure/io"

	"github.com/jinzhu/copier"
)

type ugcUsecase struct {
	BaseRepository       repository.BaseRepository
	InternshipRepository repository.InternshipRepository
	UserDataRepository   repository.UserDataRepository
	UgcRepository        repository.UgcRepository
	InterestRepository   repository.InterestRepository
	AccessUsecase        AccessUsecase
}

type UgcUsecase interface {
	List(idinterest int) (result []interface{}, err error)
}

func NewUgcUsecase(
	br repository.BaseRepository,
	ir repository.InternshipRepository,
	udr repository.UserDataRepository,
	ugc repository.UgcRepository,
	itr repository.InterestRepository,
	au AccessUsecase,
) UgcUsecase {
	return &ugcUsecase{
		br,
		ir,
		udr,
		ugc,
		itr,
		au,
	}
}

func (uu *ugcUsecase) List(idinterest int) (result []interface{}, err error) {
	Filter := map[string]interface{}{
		"interest": idinterest,
	}

	listugc, _ := uu.UgcRepository.List()
	Filterlistugc, _ := uu.UgcRepository.FindByParam(Filter)
	logger.Infof("Input Data Filterlistugc", Filterlistugc)

	if Filterlistugc != nil {
		for _, dataugc := range Filterlistugc {
			datalistugc := rsp.UgcList{}
			copier.Copy(&datalistugc, &dataugc)

			profileugc, _ := uu.UserDataRepository.Find(dataugc.IDUser)
			dataprofileugc := rsp.Ugcprofile{}
			copier.Copy(&dataprofileugc, &profileugc)

			datalistugc.Ugcprofile = append(datalistugc.Ugcprofile, dataprofileugc)
			result = append(result, datalistugc)
		}
		return result, err
	}

	for _, dataugc := range listugc {
		datalistugc := rsp.UgcList{}
		copier.Copy(&datalistugc, &dataugc)

		profileugc, _ := uu.UserDataRepository.Find(dataugc.IDUser)
		logger.Infof("Input Data profileugc", profileugc)
		dataprofileugc := rsp.Ugcprofile{}
		copier.Copy(&dataprofileugc, &profileugc)

		datalistugc.Ugcprofile = append(datalistugc.Ugcprofile, dataprofileugc)
		result = append(result, datalistugc)
	}
	return result, err
}
