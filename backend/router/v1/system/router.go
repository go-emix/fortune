package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-emix/fortune/backend/api/v1/system"
)

func Register(r gin.IRouter) {
	gr := r.Group("system")
	gr.POST("login", system.Login)
	gr.GET("tq", system.Tianqi)
	gr.GET("i18n", system.I18n)

	// 查询管理员的菜单
	gr.GET("menus", system.LoginInterceptor, system.ApiInterceptor, system.Menus)
	// 查询管理员的功能
	gr.GET("features", system.LoginInterceptor, system.ApiInterceptor, system.Features)

	// 查询所有角色
	gr.GET("roleList", system.LoginInterceptor, system.ApiInterceptor, system.RoleList)
	// 查询所有功能
	gr.GET("featureList", system.LoginInterceptor, system.ApiInterceptor, system.FeatureList)
	// 查询所有api
	gr.GET("apiList", system.LoginInterceptor, system.ApiInterceptor, system.ApiList)

	// 查询角色的功能
	gr.GET("featureListByRole", system.LoginInterceptor, system.ApiInterceptor, system.FeatureListByRole)
	// 查询角色的api
	gr.GET("apiListByRole", system.LoginInterceptor, system.ApiInterceptor, system.ApiListByRole)
	// 更新角色的功能
	gr.PUT("features", system.LoginInterceptor, system.ApiInterceptor, system.UpdateRoleFeatures)
	// 更新角色的api
	gr.PUT("apis", system.LoginInterceptor, system.ApiInterceptor, system.UpdateRoleApis)
	// 新建角色
	gr.POST("role", system.LoginInterceptor, system.ApiInterceptor, system.NewRole)
	// 删除角色
	gr.DELETE("role", system.LoginInterceptor, system.ApiInterceptor, system.DeleteRole)

}
