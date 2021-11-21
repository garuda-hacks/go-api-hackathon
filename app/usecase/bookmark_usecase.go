package usecase

import (
	rsp "github.com/garuda-hacks/go-api-hackathon/app/model/response"
	"github.com/garuda-hacks/go-api-hackathon/app/repository"
	"github.com/jinzhu/copier"
)

type bookmarkUsecase struct {
	BaseRepository        repository.BaseRepository
	BookmarkRepository    repository.BookmarkRepository
	InternRepository      repository.InternshipRepository
	ScholarshipRepository repository.ScholarRepository
	InternUsecase         InternUsecase
	ScholarUsecase        ScholarUsecase
}

type BookmarkUsecase interface {
	List() ([]interface{}, error)
}

func NewBookmarkUsecase(
	br repository.BaseRepository,
	ir repository.BookmarkRepository,
	bir repository.InternshipRepository,
	sr repository.ScholarRepository,
	udr InternUsecase,
	au ScholarUsecase,
) BookmarkUsecase {
	return &bookmarkUsecase{
		br,
		ir,
		bir,
		sr,
		udr,
		au,
	}
}

func (uu *bookmarkUsecase) List() ([]interface{}, error) {
	var result []interface{}

	listbookmark, err := uu.BookmarkRepository.List()

	for _, databookmark := range listbookmark {
		datalistbookmark := rsp.Bookmark{}
		copier.Copy(&datalistbookmark, &databookmark)

		internship, _ := uu.InternRepository.Find(databookmark.IDIntern)
		datalistinternship := rsp.BookmarkInternship{}
		copier.Copy(&datalistinternship, &internship)

		scholarship, _ := uu.ScholarshipRepository.Find(databookmark.IDScholarship)
		datalistscholarship := rsp.BookmarkScholarship{}
		copier.Copy(&datalistscholarship, &scholarship)

		datalistbookmark.Internship = append(datalistbookmark.Internship, datalistinternship)
		datalistbookmark.Scholarship = append(datalistbookmark.Scholarship, datalistscholarship)

		result = append(result, datalistbookmark)
	}
	return result, err

}
