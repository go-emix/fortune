package i18n

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	emlogrus "github.com/go-emix/emix-logrus"
	"github.com/go-emix/fortune/backend/pkg/common"
	"github.com/go-emix/fortune/backend/pkg/resp"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"io/fs"
	"path/filepath"
)

var bundle *i18n.Bundle

func Initialize(path string) error {
	bundle = i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			_, err = bundle.LoadMessageFile(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func Localize(c *gin.Context, mid string) string {
	lang := c.GetHeader("lang")
	header := c.GetHeader("Accept-Language")
	localizer := i18n.NewLocalizer(bundle, lang, header)
	message, err := localizer.LocalizeMessage(&i18n.Message{ID: mid})
	if err != nil {
		emlogrus.Error(err.Error())
		return ""
	}
	return message
}

type Error struct {
	msgId string
	c     *gin.Context
	err   error
}

func (e *Error) Error() string {
	if e.err != nil {
		return e.err.Error()
	}
	r, ok := common.Resps[e.msgId]
	if ok && r.ErrMsg != "" {
		return r.ErrMsg
	}
	return Localize(e.c, e.msgId)
}

func (e *Error) Resp() resp.Resp {
	r, ok := common.Resps[e.msgId]
	if !ok {
		r = common.Resps["system_fault"]
	}
	r.ErrMsg = e.Error()
	return r
}

func NewErr(c *gin.Context, msgId string, err error) *Error {
	return &Error{
		c:     c,
		msgId: msgId,
		err:   err,
	}
}
