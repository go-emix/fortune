package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-emix/fortune/backend/pkg/common"
	"github.com/go-emix/fortune/backend/pkg/i18n"
	"github.com/go-emix/fortune/backend/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"sort"
)

func Login(c *gin.Context, param LoginParam) (succ LoginSucc, ierr *i18n.Error) {
	var u Admin
	var err error
	err = common.DB.Model(Admin{}).Where("username=?",
		param.Username).Find(&u).Error
	if err != nil {
		ierr = i18n.NewErr(c, "", err)
		return
	}
	if u.Id == 0 {
		ierr = i18n.NewErr(c, "user_not_exist", nil)
		return
	}
	if !u.Enabled {
		ierr = i18n.NewErr(c, "user_not_enabled", nil)
		return
	}
	if u.Password == "" {
		ierr = i18n.NewErr(c, "", errors.New("password not setting"))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(param.Password))
	if err != nil {
		ierr = i18n.NewErr(c, "password_error", nil)
		return
	}
	cer := jwt.Certification{Username: u.Username, UserId: u.Id}
	token, err := jwt.GenToken(cer)
	if err != nil {
		ierr = i18n.NewErr(c, "", err)
		return
	}
	succ.Username = u.Username
	succ.Id = u.Id
	succ.Nickname = u.Nickname
	succ.Token = token
	return
}

func Menus(roles []int) (rm []Menu, err error) {
	if len(roles) == 0 {
		err = errors.New("roles not empty")
		return
	}
	root := Role{}
	err = common.DB.Where("name=?", "root").First(&root).Error
	if err != nil {
		return
	}
	if root.Id == roles[0] {
		err = common.DB.Model(Menu{}).Order("id").Find(&rm).Error
		return
	}
	for _, r := range roles {
		menus := make([]Menu, 0)
		err = common.DB.Model(Menu{}).Joins("right join role_menu on "+
			"menu.id=role_menu.menu and role_menu.role=?", r).
			Find(&menus).Error
		if err != nil {
			return
		}
		rm = append(rm, menus...)
	}
	m := make(map[int]Menu)
	mids := make([]int, 0)
	for _, menu := range rm {
		_, ok := m[menu.Id]
		if !ok {
			m[menu.Id] = menu
			mids = append(mids, menu.Id)
		}
	}
	sort.Ints(mids)
	rm = make([]Menu, 0)
	for _, mid := range mids {
		rm = append(rm, m[mid])
	}
	return
}

func Features(roles []int) (rm []Feature, err error) {
	if len(roles) == 0 {
		err = errors.New("roles not empty")
		return
	}
	root := Role{}
	err = common.DB.Where("name=?", "root").First(&root).Error
	if err != nil {
		return
	}
	defer func() {
		rm = SetMenuEntity(rm)
	}()
	if root.Id == roles[0] {
		err = common.DB.Model(Feature{}).Order("id").Find(&rm).Error
		return
	}
	for _, r := range roles {
		feas := make([]Feature, 0)
		err = common.DB.Model(Feature{}).Joins("right join role_feature on "+
			"feature.id=role_feature.feature and role_feature.role=?", r).
			Find(&feas).Error
		if err != nil {
			return
		}
		rm = append(rm, feas...)
	}
	m := make(map[int]Feature)
	fids := make([]int, 0)
	for _, fea := range rm {
		_, ok := m[fea.Id]
		if !ok {
			m[fea.Id] = fea
			fids = append(fids, fea.Id)
		}
	}
	sort.Ints(fids)
	rm = make([]Feature, 0)
	for _, fid := range fids {
		rm = append(rm, m[fid])
	}
	return
}
