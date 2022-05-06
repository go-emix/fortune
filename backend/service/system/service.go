package system

import (
	"errors"
	"github.com/go-emix/fortune/backend/pkg/common"
	"github.com/go-emix/fortune/backend/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Login(param LoginParam) (succ LoginSucc, err error) {
	var u User
	err = common.DB.Model(User{}).Where("username=?",
		param.Username).Find(&u).Error
	if err != nil {
		return
	}
	if u.Id == 0 {
		err = errors.New("user not exist")
		return
	}
	if u.Password == "" {
		err = errors.New("password not setting")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(param.Password))
	if err != nil {
		return
	}
	cer := jwt.Certification{Username: u.Username, UserId: u.Id}
	token, err := jwt.GenToken(cer)
	if err != nil {
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
