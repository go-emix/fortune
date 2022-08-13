package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-emix/fortune/backend/pkg/i18n"
	"github.com/go-emix/fortune/backend/pkg/resp"
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

func Features(c *gin.Context) {
	feas, err := system.Features(getRids(getRoles(c)))
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	resp.Data(c, feas)
}

func AdminList(c *gin.Context) {
	as, err := system.AdminList()
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	resp.Data(c, as)
}

func NewAdmin(c *gin.Context) {
	var da = struct {
		Rids []int
		system.Admin
	}{}
	err := c.ShouldBindJSON(&da)
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	err = system.NewAdmin(da.Admin, da.Rids)
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	resp.Succ(c)
}

func DeleteAdmin(c *gin.Context) {
	var da = struct {
		Id int `form:"id"`
	}{}
	err := c.ShouldBindQuery(&da)
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	err = system.DeleteAdmin(da.Id)
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	resp.Succ(c)
}

func GetAdmin(c *gin.Context) {
	var da = struct {
		Id int `form:"id"`
	}{}
	err := c.ShouldBindQuery(&da)
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	ad, err := system.GetAdmin(da.Id)
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	resp.Data(c, ad)
}
