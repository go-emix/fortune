package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-emix/fortune/backend/pkg/casbin"
	"github.com/go-emix/fortune/backend/pkg/common"
	"github.com/go-emix/fortune/backend/pkg/i18n"
	"github.com/go-emix/fortune/backend/pkg/jwt"
	"github.com/go-emix/fortune/backend/pkg/resp"
	"github.com/go-emix/fortune/backend/service/system"
)

func LoginInterceptor(c *gin.Context) {
	th := c.GetHeader("token")
	if th == "" {
		resp.Err(c, i18n.NewErr(c, "not_login", nil).Resp())
		return
	}
	token, err := jwt.ParseToken(th)
	if err != nil {
		if err == jwt.TokenExpiredErr {
			resp.Err(c, i18n.NewErr(c, "token_expired", nil).Resp())
			return
		}
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	c.Set("uid", token.UserId)
	rids := make([]int, 0)
	err = common.DB.Model(system.AdminRole{}).Joins("left join admin on "+
		"admin.id=admin_role.admin").Where("admin=?", token.UserId).
		Pluck("role", &rids).Error
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	roles := make([]system.Role, 0)
	err = common.DB.Model(system.Role{}).Where("id in (?)", rids).
		Find(&roles).Error
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	c.Set("roles", roles)
}

func ApiInterceptor(c *gin.Context) {
	roles := getRoles(c)
	if len(roles) == 0 {
		resp.Err(c, i18n.NewErr(c, "permission_denied", nil).Resp())
		return
	}
	path := c.Request.URL.Path
	method := c.Request.Method
	pass := false
	for _, role := range roles {
		pass, _ = casbin.Enforcer.Enforce(role.Name, path, method)
		if pass {
			break
		}
	}
	if !pass {
		resp.Err(c, i18n.NewErr(c, "permission_denied", nil).Resp())
	}
}
