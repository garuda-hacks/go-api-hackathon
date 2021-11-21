package router

import (
	"github.com/garuda-hacks/go-api-hackathon/app/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine, c controller.UserController) *gin.Engine {
	gr := r.Group("api")
	{
		gr.POST("/loginSecure", c.LoginUsername)
	}
	return r
}
func ProfilRouter(r *gin.Engine, c controller.ProfilController) *gin.Engine {
	gr := r.Group("api")
	{
		gr.GET("/GetUserDataWithAccountState", c.FindProfil)
	}
	return r
}
func ScholarshipRouter(r *gin.Engine, c controller.ScholarController) *gin.Engine {
	gr := r.Group("api")
	{
		gr.GET("/GetAllScholarship", c.ListScholar)
		gr.GET("/Scholarship/:id", c.FindByID)
	}
	return r
}
func InternshipRouter(r *gin.Engine, c controller.InternController) *gin.Engine {
	gr := r.Group("api")
	{
		gr.GET("/GetAllInternship", c.ListIntern)
		gr.GET("/Internship/:id", c.FindByID)
	}
	return r
}
func UgcRouter(r *gin.Engine, c controller.UgcController) *gin.Engine {
	gr := r.Group("api")
	{
		gr.GET("/getallugc", c.ListUgc)
	}
	return r
}
func MessageRouter(r *gin.Engine, c controller.MessageController) *gin.Engine {
	gr := r.Group("api")
	{
		gr.GET("/getallmessage", c.Listmessage)
	}
	return r
}
func BookmarkRouter(r *gin.Engine, c controller.BookmarkController) *gin.Engine {
	gr := r.Group("api")
	{
		gr.GET("/getallbookmark", c.Listbookmark)
	}
	return r
}
