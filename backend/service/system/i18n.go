package system

import "github.com/go-emix/fortune/backend/pkg/common"

func I18n() (en map[string]string, zh map[string]string) {
	en = make(map[string]string)
	zh = make(map[string]string)
	ns := make([]FrontI18N, 0)
	common.DB.Model(FrontI18N{}).Find(&ns)
	for _, n := range ns {
		en[n.Name] = n.En
		zh[n.Name] = n.Zh
	}
	return
}
