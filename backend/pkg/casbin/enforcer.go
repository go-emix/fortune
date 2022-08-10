package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/gorm-adapter/v3"
	"github.com/go-emix/fortune/backend/pkg/common"
	"time"
)

var rbacModel = `[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act  || r.sub == "root"`

var Enforcer *casbin.SyncedEnforcer

func Initialize() (err error) {
	ada, err := gormadapter.NewAdapterByDBUseTableName(common.DB,
		"", "api_interceptor")
	if err != nil {
		return
	}
	mod, err := model.NewModelFromString(rbacModel)
	if err != nil {
		return
	}
	Enforcer, err = casbin.NewSyncedEnforcer(mod, ada)
	if err != nil {
		return
	}
	err = Enforcer.LoadPolicy()
	if err != nil {
		return
	}
	Enforcer.StartAutoLoadPolicy(10 * time.Second)
	Enforcer.EnableAutoSave(true)
	return
}
