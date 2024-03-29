package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-emix/fortune/backend/pkg/common"
	"github.com/go-emix/fortune/backend/pkg/i18n"
	"github.com/go-emix/fortune/backend/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"sort"
)

func Login(c *gin.Context, param LoginParam) (succ LoginSucc, ierr *i18n.Error) {
	var u Admin
	var err error
	err = common.DB.Model(Admin{}).Where("username=?",
		param.Username).First(&u).Error
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

func Menus(roles []int) (rs []Menu, err error) {
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
		err = common.DB.Model(Menu{}).Order("id").Find(&rs).Error
		return
	}
	for _, r := range roles {
		menus := make([]Menu, 0)
		err = common.DB.Model(Menu{}).Joins("join role_menu on "+
			"menu.id=role_menu.menu and role_menu.role=?", r).
			Find(&menus).Error
		if err != nil {
			return
		}
		rs = append(rs, menus...)
	}
	m := make(map[int]Menu)
	mids := make([]int, 0)
	for _, menu := range rs {
		_, ok := m[menu.Id]
		if !ok {
			m[menu.Id] = menu
			mids = append(mids, menu.Id)
		}
	}
	sort.Ints(mids)
	rs = make([]Menu, 0)
	for _, mid := range mids {
		rs = append(rs, m[mid])
	}
	return
}

func Features(roles []int) (rs []Feature, err error) {
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
		rs = SetMenuEntity(rs)
	}()
	if root.Id == roles[0] {
		err = common.DB.Model(Feature{}).Order("id").Find(&rs).Error
		return
	}
	for _, r := range roles {
		feas := make([]Feature, 0)
		err = common.DB.Model(Feature{}).Joins("join role_feature on "+
			"feature.id=role_feature.feature and role_feature.role=?", r).
			Find(&feas).Error
		if err != nil {
			return
		}
		rs = append(rs, feas...)
	}
	m := make(map[int]Feature)
	fids := make([]int, 0)
	for _, fea := range rs {
		_, ok := m[fea.Id]
		if !ok {
			m[fea.Id] = fea
			fids = append(fids, fea.Id)
		}
	}
	sort.Ints(fids)
	rs = make([]Feature, 0)
	for _, fid := range fids {
		rs = append(rs, m[fid])
	}
	return
}

func Apis(roles []int) (rs []Api, err error) {
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
		err = common.DB.Model(Api{}).Order("id").Find(&rs).Error
		return
	}
	for _, r := range roles {
		as := make([]Api, 0)
		err = common.DB.Model(Api{}).Joins("join role_api on "+
			"api.id=role_api.api and role_api.role=?", r).
			Find(&as).Error
		if err != nil {
			return
		}
		rs = append(rs, as...)
	}
	m := make(map[int]Api)
	aids := make([]int, 0)
	for _, a := range rs {
		_, ok := m[a.Id]
		if !ok {
			m[a.Id] = a
			aids = append(aids, a.Id)
		}
	}
	sort.Ints(aids)
	rs = make([]Api, 0)
	for _, id := range aids {
		rs = append(rs, m[id])
	}
	return
}

func AdminList() (as []Admin, err error) {
	err = common.DB.Model(Admin{}).Find(&as).Error
	return
}

func GetAdmin(aid int) (ad Admin, err error) {
	ad.Id = aid
	err = common.DB.First(&ad).Error
	if err != nil {
		return
	}
	ad.Roles = make([]Role, 0)
	err = common.DB.Model(Role{}).Joins("join admin_role on "+
		"admin_role.role=role.id and admin_role.admin=?", ad.Id).
		Find(&ad.Roles).Error
	return
}

func NewAdmin(ad Admin, rids []int) (err error) {
	if len(rids) == 0 {
		return errors.New("rids not be empty")
	}
	var c int64
	common.DB.Model(Admin{}).Where("username=?", ad.Username).Count(&c)
	if c != 0 {
		err = errors.New("admin username already exist")
		return
	}
	err = common.DB.Transaction(func(tx *gorm.DB) (err error) {
		password, err := bcrypt.GenerateFromPassword([]byte(ad.Password), bcrypt.DefaultCost)
		if err != nil {
			return
		}
		ad.Password = string(password)
		err = common.DB.Create(&ad).Error
		if err != nil {
			return
		}
		for _, rid := range rids {
			err = common.DB.Create(&AdminRole{Admin: ad.Id, Role: rid}).Error
			if err != nil {
				return
			}
		}
		return
	})
	return
}

func DeleteAdmin(aid int) (err error) {
	var ad Admin
	ad.Id = aid
	err = common.DB.First(&ad).Error
	if err != nil {
		return
	}
	err = common.DB.Transaction(func(tx *gorm.DB) (err error) {
		err = common.DB.Unscoped().Delete(&ad).Error
		if err != nil {
			return
		}
		err = common.DB.Where("admin=?", aid).Unscoped().Delete(AdminRole{}).Error
		return
	})
	return
}

func UpdateAdmin(ad Admin, rids []int) (err error) {
	if len(rids) == 0 {
		return errors.New("rids not be empty")
	}
	if ad.Password != "" {
		var pass []byte
		pass, err = bcrypt.GenerateFromPassword([]byte(ad.Password), bcrypt.DefaultCost)
		if err != nil {
			return
		}
		ad.Password = string(pass)
	}
	err = common.DB.Transaction(func(tx *gorm.DB) (err error) {
		if ad.Password == "" {
			err = common.DB.Omit("password").Save(&ad).Error
		} else {
			err = common.DB.Save(&ad).Error
		}
		if err != nil {
			return
		}
		err = common.DB.Where("admin=?", ad.Id).Unscoped().Delete(AdminRole{}).Error
		if err != nil {
			return
		}
		for _, rid := range rids {
			err = common.DB.Create(&AdminRole{Admin: ad.Id, Role: rid}).Error
			if err != nil {
				return
			}
		}
		return
	})
	return
}
