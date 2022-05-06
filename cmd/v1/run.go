package v1

import (
	"github.com/go-emix/fortune/backend/pkg/common"
	"github.com/go-emix/fortune/backend/pkg/jwt"
	"github.com/go-emix/fortune/backend/server"
	"github.com/go-emix/fortune/backend/service"
	"github.com/go-emix/fortune/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
)

func run() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run server,default port " + strconv.Itoa(config.ServerC.Port),
	}
	cmd.Flags().IntVarP(&config.ServerC.Port, "port", "p",
		0, "server port")
	cmd.Flags().StringVarP(&cf, "conf", "c",
		"", "config file")
	cmd.Run = func(cmd *cobra.Command, _ []string) {
		err := loadConfig()
		if err != nil {
			cmd.Println(err.Error())
			return
		}
		config.ExeSql()
		common.Resps = config.AppC.Resps
		common.DB = config.GetDb()
		jwt.Initialize(config.AppC.Jwt.SignKey, config.AppC.Jwt.Expire)
		err = service.Migrate()
		if err != nil {
			cmd.Println(err.Error())
			return
		}
		server.SetMode(string(config.AppC.Mode))
		backend := server.NewServer(config.ServerC.Port, config.FrontendC.Dist)
		backend.Run()
	}
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		panic(err)
	}
	return cmd
}
