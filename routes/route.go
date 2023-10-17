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
		admin.PUT("/admin/updateUser/:id", rtr.santriService.UpdateUserForAdmin)
	}

	User := r.Group("/", middleware.UserMiddleware())
	{
		User.GET("/profile", rtr.santriService.GetProfile)
		User.POST("/createSantri", rtr.santriService.CreateSantri)
		User.GET("/searchSantri", rtr.santriService.SearchHandler)
		User.PUT("/updateUser/:id", rtr.santriService.UpdateUserProfile)
		User.DELETE("/deleteUser/:id", rtr.santriService.DeleteUser)
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

	return r
}
