package registry

import (
	"github.com/garuda-hacks/go-api-hackathon/app/repository"
	"github.com/garuda-hacks/go-api-hackathon/app/usecase"
)

func (r *registry) NewProfilUsecase() usecase.ProfilUsecase {
	return usecase.NewProfilUsecase(
		repository.NewBaseRepository(r.db),
		repository.NewUserDataRepository(repository.NewBaseRepository(r.db)),
		repository.NewUgcRepository(repository.NewBaseRepository(r.db)),
		repository.NewInterestRepository(repository.NewBaseRepository(r.db)),
		usecase.NewAccessUsecase(
			repository.NewBaseRepository(r.db),
			repository.NewUserAccessRepository(repository.NewBaseRepository(r.db)),
		),
	)
}
func (r *registry) NewAccessUsecase() usecase.AccessUsecase {
	return usecase.NewAccessUsecase(
		repository.NewBaseRepository(r.db),
		repository.NewUserAccessRepository(repository.NewBaseRepository(r.db)),
	)
}
func (r *registry) NewUserUsecase() usecase.UserUsecase {
	return usecase.NewUserUsecase(
		repository.NewBaseRepository(r.db),
		repository.NewUserRepository(repository.NewBaseRepository(r.db)),
		usecase.NewAccessUsecase(
			repository.NewBaseRepository(r.db),
			repository.NewUserAccessRepository(repository.NewBaseRepository(r.db)),
		),
	)
}
func (r *registry) NewScholarUsecase() usecase.ScholarUsecase {
	return usecase.NewScholarUsecase(
		repository.NewBaseRepository(r.db),
		repository.NewScholarRepository(repository.NewBaseRepository(r.db)),
		repository.NewUserDataRepository(repository.NewBaseRepository(r.db)),
	)
}
func (r *registry) NewInternUsecase() usecase.InternUsecase {
	return usecase.NewInternUsecase(
		repository.NewBaseRepository(r.db),
		repository.NewInternshipRepository(repository.NewBaseRepository(r.db)),
		repository.NewUserDataRepository(repository.NewBaseRepository(r.db)),
	)
}
func (r *registry) NewUgcUsecase() usecase.UgcUsecase {
	return usecase.NewUgcUsecase(
		repository.NewBaseRepository(r.db),
		repository.NewInternshipRepository(repository.NewBaseRepository(r.db)),
		repository.NewUserDataRepository(repository.NewBaseRepository(r.db)),
		repository.NewUgcRepository(repository.NewBaseRepository(r.db)),
		repository.NewInterestRepository(repository.NewBaseRepository(r.db)),
		usecase.NewAccessUsecase(
			repository.NewBaseRepository(r.db),
			repository.NewUserAccessRepository(repository.NewBaseRepository(r.db)),
		),
	)
}
func (r *registry) NewMessageUsecase() usecase.MessageUsecase {
	return usecase.NewMessageUsecase(
		repository.NewBaseRepository(r.db),
		repository.NewMessageRepository(repository.NewBaseRepository(r.db)),
		usecase.NewAccessUsecase(
			repository.NewBaseRepository(r.db),
			repository.NewUserAccessRepository(repository.NewBaseRepository(r.db)),
		),
	)
}
func (r *registry) NewBookmarkUsecase() usecase.BookmarkUsecase {
	return usecase.NewBookmarkUsecase(
		repository.NewBaseRepository(r.db),
		repository.NewBookmarkRepository(repository.NewBaseRepository(r.db)),
		repository.NewInternshipRepository(repository.NewBaseRepository(r.db)),
		repository.NewScholarRepository(repository.NewBaseRepository(r.db)),
		usecase.NewInternUsecase(
			repository.NewBaseRepository(r.db),
			repository.NewInternshipRepository(repository.NewBaseRepository(r.db)),
			repository.NewUserDataRepository(repository.NewBaseRepository(r.db)),
		),
		usecase.NewScholarUsecase(
			repository.NewBaseRepository(r.db),
			repository.NewScholarRepository(repository.NewBaseRepository(r.db)),
			repository.NewUserDataRepository(repository.NewBaseRepository(r.db)),
		),
	)
}
