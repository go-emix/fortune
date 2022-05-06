package v1

import (
	"errors"
	emlogrus "github.com/go-emix/emix-logrus"
	"github.com/go-emix/fortune/backend/pkg/resp"
	"github.com/go-emix/fortune/config"
	"github.com/go-emix/fortune/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func Exe() {
	command := &cobra.Command{
		Use:     version.Name,
		Short:   version.ShortName,
		Version: version.Version,
	}
	command.SetFlagErrorFunc(func(_ *cobra.Command, _ error) error {
		command.Println("unknown command,use -h or --help")
		return nil
	})
	command.AddCommand(run())
	err := command.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var cf string

func loadConfig() error {
	if cf != "" {
		viper.SetConfigFile(cf)
	} else {
		viper.SetConfigType("yml")
		viper.SetConfigName("app")
		viper.AddConfigPath("./")
		viper.AddConfigPath("conf")
	}
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	database := viper.GetString("db.database")
	if database != "" {
		config.DbC.Database = database
	}
	password := viper.GetString("db.password")
	if password != "" {
		config.DbC.Password = password
	}
	dbIp := viper.GetString("db.ip")
	if dbIp != "" {
		config.DbC.Ip = dbIp
	}
	dbPort := viper.GetInt("db.port")
	if dbPort != 0 {
		config.DbC.Port = dbPort
	}
	scripts := viper.GetStringSlice("db.scripts")
	if scripts != nil {
		config.DbC.Scripts = scripts
	}
	port := viper.GetInt("server.port")
	if port != 0 {
		config.ServerC.Port = port
	}
	frdist := viper.GetString("frontend.dist")
	if frdist != "" {
		config.FrontendC.Dist = frdist
	}
	mo := viper.GetString("app.mode")
	if mo != "" {
		m := config.Mode(mo)
		if m != config.Dev {
			if m != config.Prod {
				return errors.New("unknown app mode " + string(m))
			}
		}
		config.AppC.Mode = m
	}
	get := viper.Get("logs")
	ls, ok := get.([]interface{})
	if ok {
		lcs := make([]emlogrus.LogConfig, 0)
		for _, l := range ls {
			m := l.(map[interface{}]interface{})
			lc := emlogrus.LogConfig{
				OutType:  emlogrus.ConsoleOut,
				OutDir:   "log",
				Format:   emlogrus.TextLog,
				Disabled: false,
			}
			if m["level"] == nil {
				return errors.New("level not empty")
			} else {
				lc.Level = emlogrus.LogLevel(m["level"].(string))
			}
			if ot := m["outType"]; ot != nil {
				lc.OutType = emlogrus.LogOutType(ot.(string))
			}
			if sl := m["singleLevel"]; sl != nil {
				lc.SingleLevel = sl.(bool)
			}
			if ot := m["outDir"]; ot != nil {
				lc.OutDir = ot.(string)
			}
			if fr := m["format"]; fr != nil {
				lc.Format = emlogrus.LogFormat(fr.(string))
			}
			if di := m["disabled"]; di != nil {
				lc.Disabled = di.(bool)
			}
			if ma := m["maxAge"]; ma != nil {
				lc.MaxAge = ma.(int)
			}
			if mc := m["maxCount"]; mc != nil {
				lc.MaxCount = uint(mc.(int))
			}
			lcs = append(lcs, lc)
		}
		emlogrus.AfterInit("", lcs...).Setup()
	}
	resps := viper.Get("app.resps")
	rs, ok := resps.([]interface{})
	if ok {
		rps := make(map[string]resp.Resp, 0)
		for _, r := range rs {
			m := r.(map[interface{}]interface{})
			rp := resp.Resp{}
			if m["errcode"] == nil {
				return errors.New("errcode not empty")
			} else {
				rp.ErrCode = m["errcode"].(int)
			}
			if m["errmsg"] == nil {
				return errors.New("errmsg not empty")
			} else {
				rp.ErrMsg = m["errmsg"].(string)
			}
			if m["name"] == nil {
				return errors.New("name not empty")
			} else {
				rps[m["name"].(string)] = rp
			}
		}
		config.AppC.Resps = rps
	}
	expire := viper.GetInt("app.jwt.expire")
	if expire != 0 {
		config.AppC.Jwt.Expire = expire
	}
	signKey := viper.GetString("app.jwt.key")
	if signKey != "" {
		config.AppC.Jwt.SignKey = signKey
	}
	return nil
}
