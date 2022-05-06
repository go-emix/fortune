package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-emix/fortune/backend/router/v1/system"
)

func Register(r gin.IRouter) {
	v1 := r.Group("v1")
	system.Register(v1)
}
