package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-emix/fortune/backend/pkg/i18n"
	"github.com/go-emix/fortune/backend/pkg/resp"
	"github.com/go-emix/fortune/backend/pkg/tianqi"
	"github.com/go-emix/fortune/backend/service/system"
)

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

func FeatureList(c *gin.Context) {
	feas, err := system.FeatureList()
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	resp.Data(c, feas)
}

func FeatureListByRole(c *gin.Context) {
	var da = struct {
		Role int `form:"role"`
	}{}
	err := c.ShouldBindQuery(&da)
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	if da.Role == 0 {
		resp.Err(c, i18n.NewErr(c, "", errors.New("query role miss")).Resp())
		return
	}
	feas, err := system.Features([]int{da.Role})
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	resp.Data(c, feas)
}
