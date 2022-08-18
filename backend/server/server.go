package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	emlogrus "github.com/go-emix/emix-logrus"
	"github.com/go-emix/fortune/backend/pkg/common"
	"github.com/go-emix/fortune/backend/router"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

type Server struct {
	eng      *gin.Engine
	port     int
	killChan chan os.Signal
}

func setMode() {
	switch common.RunMode {
	case common.Dev:
		gin.SetMode(gin.DebugMode)
	case common.Prod:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}

func NewServer(port int, frontDist string) *Server {
	setMode()
	ser := &Server{port: port, eng: gin.New()}
	ser.killChan = make(chan os.Signal)
	ser.eng.Use(gin.Recovery())
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	corsConf.AddAllowHeaders("token", "lang")
	ser.eng.Use(cors.New(corsConf))
	if frontDist != "" {
		ser.eng.GET("", func(c *gin.Context) {
			file, err := os.ReadFile(frontDist + "/index.html")
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError,
					"server fault")
				return
			}
			c.Data(http.StatusOK, "text/html;charset=utf-8", file)
		})
		ser.eng.Static("assets", frontDist+"/assets")
		ser.eng.GET("favicon.ico", func(c *gin.Context) {
			c.File(frontDist + "/favicon.ico")
		})
		api := ser.eng.Group("api")
		router.Register(api)
	} else {
		router.Register(ser.eng)
	}
	return ser
}

func (r *Server) Run() {
	go func() {
		signal.Notify(r.killChan, os.Interrupt)
		emlogrus.Info("server start")
		err := r.eng.Run(":" + strconv.Itoa(r.port))
		if err != nil {
			emlogrus.Error(err.Error())
			r.killChan <- os.Interrupt
		}
	}()
	for {
		sig := <-r.killChan
		if sig == os.Interrupt {
			emlogrus.Info("server stop")
			os.Exit(1)
		}
	}
}
