package registry

import "github.com/garuda-hacks/go-api-hackathon/app/controller"

func (r *registry) NewProfilController() controller.ProfilController {
	return controller.NewProfilController(r.NewProfilUsecase())
}
func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserUsecase())
}
func (r *registry) NewScholarController() controller.ScholarController {
	return controller.NewScholarController(r.NewScholarUsecase())
}
func (r *registry) NewInternController() controller.InternController {
	return controller.NewInternController(r.NewInternUsecase())
}
func (r *registry) NewUgcController() controller.UgcController {
	return controller.NewUgcController(r.NewUgcUsecase())
}
func (r *registry) NewMessageController() controller.MessageController {
	return controller.NewMessageController(r.NewMessageUsecase())
}
func (r *registry) NewBookmarkController() controller.BookmarkController {
	return controller.NewBookmarkController(r.NewBookmarkUsecase())
}
