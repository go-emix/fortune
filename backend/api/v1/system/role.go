package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-emix/fortune/backend/pkg/i18n"
	"github.com/go-emix/fortune/backend/pkg/resp"
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

func FeatureList(c *gin.Context) {
	feas, err := system.FeatureList()
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	resp.Data(c, feas)
}

func ApiList(c *gin.Context) {
	aps, err := system.ApiList()
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	resp.Data(c, aps)
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
	rs, err := system.Features([]int{da.Role})
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	resp.Data(c, rs)
}

func ApiListByRole(c *gin.Context) {
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
	rs, err := system.Apis([]int{da.Role})
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	resp.Data(c, rs)
}

func UpdateRoleFeatures(c *gin.Context) {
	var da = struct {
		Rid  int
		Fids []int
	}{}
	err := c.ShouldBindJSON(&da)
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	if da.Rid == 0 {
		resp.Err(c, i18n.NewErr(c, "", errors.New("body rid miss")).Resp())
		return
	}
	err = system.UpdateRoleFeatures(da.Rid, da.Fids)
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	resp.Succ(c)
}

func UpdateRoleApis(c *gin.Context) {
	var da = struct {
		Rid  int
		Aids []int
	}{}
	err := c.ShouldBindJSON(&da)
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	if da.Rid == 0 {
		resp.Err(c, i18n.NewErr(c, "", errors.New("body rid miss")).Resp())
		return
	}
	err = system.UpdateRoleApis(da.Rid, da.Aids)
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	resp.Succ(c)
}

func NewRole(c *gin.Context) {
	var da = struct {
		Name string
	}{}
	err := c.ShouldBindJSON(&da)
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	if da.Name == "" {
		resp.Err(c, i18n.NewErr(c, "", errors.New("body name miss")).Resp())
		return
	}
	err = system.NewRole(da.Name)
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	resp.Succ(c)
}

func DeleteRole(c *gin.Context) {
	var da = struct {
		Id int `form:"id"`
	}{}
	err := c.ShouldBindQuery(&da)
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	if da.Id == 0 {
		resp.Err(c, i18n.NewErr(c, "", errors.New("query id miss")).Resp())
		return
	}
	err = system.DeleteRole(da.Id)
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	resp.Succ(c)
}
