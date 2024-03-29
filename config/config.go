package config

import (
	"github.com/go-emix/fortune/backend/pkg/resp"
	"github.com/go-emix/fortune/version"
	"github.com/go-emix/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
)

type Db struct {
	Ip       string
	Port     int
	Database string
	Password string
	Scripts  []string
}

type Server struct {
	Port int
}

type Frontend struct {
	Dist string
}

type App struct {
	Mode  string
	Resps map[string]resp.Resp
	Jwt   Jwt
	I18n  string
}

type Jwt struct {
	SignKey string
	Expire  int // second
}

var ServerC Server
var DbC Db
var AppC App
var FrontendC Frontend

func init() {
	DbC.Ip = "127.0.0.1"
	DbC.Port = 3306
	DbC.Database = version.Name
	DbC.Password = "123456"
	DbC.Scripts = []string{}
	ServerC.Port = 3678
	AppC.Mode = "dev"
	AppC.Resps = make(map[string]resp.Resp)
	// init resps
	errcode := 1000
	AppC.Resps["system_fault"] = resp.Resp{ErrCode: errcode}
	errcode++
	AppC.Resps["user_not_exist"] = resp.Resp{ErrCode: errcode}
	errcode++
	AppC.Resps["password_error"] = resp.Resp{ErrCode: errcode}
	errcode++
	AppC.Resps["user_not_enabled"] = resp.Resp{ErrCode: errcode}
	errcode++
	AppC.Resps["token_expired"] = resp.Resp{ErrCode: errcode}
	errcode++
	AppC.Resps["permission_denied"] = resp.Resp{ErrCode: errcode}
	AppC.Jwt.SignKey = version.Name
	AppC.Jwt.Expire = 3600
	AppC.I18n = "conf/i18n"
}

func GetDb() *gorm.DB {
	dsn := "root:" + DbC.Password + "@tcp(" + DbC.Ip + ":" + strconv.Itoa(DbC.Port) + ")/" +
		DbC.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func ExeSql() {
	db := GetDb()
	defer func() {
		d, err := db.DB()
		if err != nil {
			panic(err)
		}
		err = d.Close()
		if err != nil {
			panic(err)
		}
	}()
	for _, sf := range DbC.Scripts {
		if !utils.FileIsExist(sf) {
			continue
		}
		file, err := os.ReadFile(sf)
		if err != nil {
			panic(err)
		}
		sql := string(file)
		sqlList := strings.Split(sql, ";")
		for _, sql := range sqlList {
			spa := strings.TrimSpace(sql)
			if spa == "" {
				continue
			}
			err = db.Exec(spa + ";").Error
			if err != nil {
				panic(err)
			}
		}
	}
}
