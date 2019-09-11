package router

import (
	"elib_unikom/controller"

	"github.com/gin-gonic/gin"
	// "goadmin_starter_pack/middleware"
)

func RouterDashboard(r *gin.Engine) {
	rDashboard := r.Group("/dashboard")
	// rDashboard.Use(middleware.Auth)
	rDashboard.GET("/", controller.Dashboard)

}
