package v1

import (
	"github.com/gin-gonic/gin"
	emlogrus "github.com/go-emix/emix-logrus"
	"github.com/go-emix/fortune/backend/api/v1/system"
)

func LogInterceptor(c *gin.Context) {
	ur := c.Request.URL.String()
	c.Next()
	uid := system.GetUid(c)
	emlogrus.Infof("uid:%d,url:%s,ip:%s", uid, ur, c.ClientIP())
}
