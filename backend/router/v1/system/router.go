package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-emix/fortune/backend/api/v1/system"
)

func Register(r gin.IRouter) {
	gr := r.Group("system")
	gr.POST("login", system.Login)
	gr.GET("tq", system.Tianqi)
	gr.GET("menus", system.LoginInterceptor, system.ApiInterceptor, system.Menus)
	gr.GET("features", system.LoginInterceptor, system.ApiInterceptor, system.Features)
	gr.GET("roleList", system.LoginInterceptor, system.ApiInterceptor, system.RoleList)
}
