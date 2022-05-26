package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	ErrCode int         `json:"errcode"`
	ErrMsg  string      `json:"errmsg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Succ(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, Resp{
		ErrCode: 0,
		ErrMsg:  "success",
	})
}

func Err(c *gin.Context, resp Resp) {
	c.AbortWithStatusJSON(http.StatusOK, resp)
}

func Data(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, Resp{
		ErrCode: 0,
		Data:    data,
	})
}
