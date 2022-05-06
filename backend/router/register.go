package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/go-emix/fortune/backend/router/v1"
)

func Register(r gin.IRouter) {
	v1.Register(r)
}
