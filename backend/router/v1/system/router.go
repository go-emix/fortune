package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-emix/fortune/backend/api/v1/system"
)

func Register(r gin.IRouter) {
	gr := r.Group("system")
	gr.POST("login", system.Login)
	gr.GET("tq", system.Tianqi)
}
