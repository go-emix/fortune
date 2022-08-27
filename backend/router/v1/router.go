package v1

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/go-emix/fortune/backend/api/v1"
	"github.com/go-emix/fortune/backend/router/v1/system"
)

func Register(r gin.IRouter) {
	v1g := r.Group("v1")
	v1g.Use(v1.LogInterceptor)
	system.Register(v1g)
}
