package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-emix/fortune/backend/pkg/common"
	"github.com/go-emix/fortune/backend/pkg/i18n"
	"github.com/go-emix/fortune/backend/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context, param LoginParam) (succ LoginSucc, ierr *i18n.Error) {
	var u User
	var err error
	err = common.DB.Model(User{}).Where("username=?",
		param.Username).Find(&u).Error
	if err != nil {
		ierr = i18n.NewErr(c, "", err)
		return
	}
	if u.Id == 0 {
		ierr = i18n.NewErr(c, "user_not_exist", nil)
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

func Migrate() error {
	err := common.DB.AutoMigrate(User{})
	if err != nil {
		return err
	}
	ad := User{
		Username: "ad",
		Nickname: "管理员",
	}
	password, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	ad.Password = string(password)
	err = common.DB.FirstOrCreate(&ad, User{Username: ad.Username}).Error
	return err
}
