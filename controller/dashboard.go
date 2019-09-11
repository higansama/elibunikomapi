package controller

import (
	"elib_unikom/lib"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	js := []string{
		"../asset/js/myjs/dashboard",
		"../asset/plugin/tagsinput/bootstrap-tagsinput",
	}
	css := []string{
		"../asset/plugin/tagsinput/bootstrap-tagsinput",
	}

	lib.Pages(c, 200, nil, "Dashboard Index", "dashboard/index", lib.MapKosong, js, css)

}
