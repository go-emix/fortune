package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-emix/fortune/backend/pkg/i18n"
	"github.com/go-emix/fortune/backend/pkg/resp"
	"github.com/go-emix/fortune/backend/pkg/tianqi"
	"github.com/go-emix/fortune/backend/service/system"
)

func Login(c *gin.Context) {
	var pa system.LoginParam
	err := c.ShouldBindJSON(&pa)
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	succ, ierr := system.Login(c, pa)
	if ierr != nil {
		resp.Err(c, ierr.Resp())
		return
	}
	resp.Data(c, succ)
}

func Menus(c *gin.Context) {
	menus, err := system.Menus(getRids(getRoles(c)))
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	resp.Data(c, menus)
}

func RoleList(c *gin.Context) {
	menus, err := system.RoleList()
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	resp.Data(c, menus)
}

func Tianqi(c *gin.Context) {
	temp, err := tianqi.GetTemp()
	if err != nil {
		resp.Err(c, resp.Resp{
			ErrCode: 2001,
			ErrMsg:  err.Error(),
		})
		return
	}
	resp.Data(c, temp)
}

func Features(c *gin.Context) {
	feas, err := system.Features(getRids(getRoles(c)))
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	resp.Data(c, feas)
}
