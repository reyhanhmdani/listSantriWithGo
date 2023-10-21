package routes

import (
	"github.com/gin-gonic/gin"
	"project1/middleware"
	"project1/service"
)

type Route struct {
	santriService *service.Handler
}

func NewRoute(santriService *service.Handler) *Route {
	return &Route{
		santriService: santriService,
	}
}

func (rtr *Route) RouteInit() *gin.Engine {

	r := gin.Default()
	r.Use(middleware.RecoveryMiddleware(), middleware.Logger())

	admin := r.Group("/", middleware.AdminMiddleware())
	{
		admin.GET("/admin/allUsers", rtr.santriService.ViewAllUsers)
		admin.PATCH("/admin/updateUser/:id", rtr.santriService.UpdateUserForAdmin)
	}

	User := r.Group("/", middleware.UserMiddleware())
	{
		User.GET("/profile", rtr.santriService.GetProfile)
		User.POST("/createSantri", rtr.santriService.CreateSantri)
		User.GET("/searchSantri", rtr.santriService.SearchHandler)
		//User.PUT("/updateUsername/:id", rtr.santriService.UpdateUserForAdmin)
		User.PATCH("/updateUsername/:id", rtr.santriService.UpdateUserUsername)
		User.PATCH("/updatePassword/:id", rtr.santriService.UpdateUserPassword)
		User.DELETE("/deleteUser/:id", rtr.santriService.DeleteUser)

		// upload
		User.POST("/uploadLocal/:id", rtr.santriService.UploadFileLocal)
	}

	r.POST("/register", rtr.santriService.Register)
	r.POST("/login", rtr.santriService.Login)
	r.GET("/allSantri", rtr.santriService.ViewAllSantri)
	r.GET("/getSantri/:id", rtr.santriService.GetSantriByID)

	// References Data
	r.GET("/getReferencesData", rtr.santriService.GetReferenceData)
	r.GET("/getJurusanList", rtr.santriService.GetJurusanList)
	r.GET("/getMinatList", rtr.santriService.GetMinatList)
	r.GET("/getStatusList", rtr.santriService.GetStatusList)

	// forgot password
	r.POST("/forgot_password", rtr.santriService.ForgotPassword)
	r.PATCH("/reset_password/:token", rtr.santriService.ResetPassword)
	return r
}
