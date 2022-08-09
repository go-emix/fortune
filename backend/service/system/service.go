package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	emlogrus "github.com/go-emix/emix-logrus"
	"github.com/go-emix/fortune/backend/pkg/common"
	"github.com/go-emix/fortune/backend/pkg/i18n"
	"github.com/go-emix/fortune/backend/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"sort"
)

func Login(c *gin.Context, param LoginParam) (succ LoginSucc, ierr *i18n.Error) {
	var u Admin
	var err error
	err = common.DB.Model(Admin{}).Where("username=?",
		param.Username).Find(&u).Error
	if err != nil {
		ierr = i18n.NewErr(c, "", err)
		return
	}
	if u.Id == 0 {
		ierr = i18n.NewErr(c, "user_not_exist", nil)
		return
	}
	if !u.Enabled {
		ierr = i18n.NewErr(c, "user_not_enabled", nil)
		return
	}
	if u.Password == "" {
		ierr = i18n.NewErr(c, "", errors.New("password not setting"))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(param.Password))
	if err != nil {
		ierr = i18n.NewErr(c, "password_error", nil)
		return
	}
	cer := jwt.Certification{Username: u.Username, UserId: u.Id}
	token, err := jwt.GenToken(cer)
	if err != nil {
		ierr = i18n.NewErr(c, "", err)
		return
	}
	succ.Username = u.Username
	succ.Id = u.Id
	succ.Nickname = u.Nickname
	succ.Token = token
	return
}

func Menus(roles []int) (rm []Menu, err error) {
	if len(roles) == 0 {
		err = errors.New("roles not empty")
		return
	}
	root := Role{}
	err = common.DB.Where("name=?", "root").First(&root).Error
	if err != nil {
		return
	}
	if root.Id == roles[0] {
		err = common.DB.Model(Menu{}).Order("id").Find(&rm).Error
		return
	}
	for _, r := range roles {
		menus := make([]Menu, 0)
		err = common.DB.Model(Menu{}).Joins("right join role_menu on "+
			"menu.id=role_menu.menu and role_menu.role=?", r).
			Find(&menus).Error
		if err != nil {
			return
		}
		rm = append(rm, menus...)
	}
	m := make(map[int]Menu)
	mids := make([]int, 0)
	for _, menu := range rm {
		_, ok := m[menu.Id]
		if !ok {
			m[menu.Id] = menu
			mids = append(mids, menu.Id)
		}
	}
	sort.Ints(mids)
	rm = make([]Menu, 0)
	for _, mid := range mids {
		rm = append(rm, m[mid])
	}
	return
}

func RoleList() (rs []Role, err error) {
	err = common.DB.Model(Role{}).Find(&rs).Error
	return
}

func Migrate() (err error) {
	// admin init
	emlogrus.Info("migrate admin init")
	err = common.DB.AutoMigrate(Admin{})
	if err != nil {
		return
	}
	ad := Admin{
		Username: "ad",
		Nickname: "管理员",
	}
	password, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	ad.Password = string(password)
	common.DB.Model(Admin{}).Where("username=?", ad.Username).Find(&ad)
	if common.RunMode == common.Dev {
		ad.Enabled = true
	} else {
		ad.Enabled = false
	}
	if ad.Id == 0 {
		err = common.DB.Create(&ad).Error
		if err != nil {
			return
		}
	} else {
		err = common.DB.Save(&ad).Error
		if err != nil {
			return
		}
	}
	// menu init
	emlogrus.Info("migrate menu init")
	err = common.DB.AutoMigrate(Menu{})
	if err != nil {
		return
	}
	login := "login"
	dashboard := "dashboard"
	dashboardMenu := Menu{
		Name:      dashboard,
		Path:      "/" + dashboard,
		Component: dashboard,
		Auth:      login,
	}
	create := common.DB.FirstOrCreate(&dashboardMenu, Menu{Name: dashboard})
	if dashboardMenu.Id == 0 {
		return create.Error
	}
	system := "system"
	systemMenu := Menu{
		Name: system,
	}
	create = common.DB.FirstOrCreate(&systemMenu, Menu{Name: system})
	if systemMenu.Id == 0 {
		return create.Error
	}
	admin := "admin"
	adminMenu := Menu{
		Name:      admin,
		Path:      "/" + admin,
		Component: admin,
		Auth:      admin,
		Parent:    systemMenu.Id,
	}
	create = common.DB.FirstOrCreate(&adminMenu, Menu{Name: admin})
	if adminMenu.Id == 0 {
		return create.Error
	}
	role := "role"
	roleMenu := Menu{
		Name:      role,
		Path:      "/" + role,
		Component: role,
		Auth:      role,
		Parent:    systemMenu.Id,
	}
	create = common.DB.FirstOrCreate(&roleMenu, Menu{Name: role})
	if roleMenu.Id == 0 {
		return create.Error
	}
	// role init
	emlogrus.Info("migrate role init")
	err = common.DB.AutoMigrate(Role{})
	if err != nil {
		return
	}
	root := "root"
	rootRole := Role{Name: "root"}
	create = common.DB.FirstOrCreate(&rootRole, Role{Name: root})
	if rootRole.Id == 0 {
		return create.Error
	}
	// role_menu init
	emlogrus.Info("migrate role_menu init")
	err = common.DB.AutoMigrate(RoleMenu{})
	if err != nil {
		return
	}
	// admin_role init
	emlogrus.Info("migrate admin_role init")
	err = common.DB.AutoMigrate(AdminRole{})
	if err != nil {
		return
	}
	userRootRole := AdminRole{Admin: ad.Id,
		Role: rootRole.Id}
	create = common.DB.FirstOrCreate(&userRootRole, userRootRole)
	if userRootRole.Id == 0 {
		return create.Error
	}
	// api init
	emlogrus.Info("migrate api init")
	err = common.DB.AutoMigrate(Api{})
	if err != nil {
		return
	}
	// role_api init
	emlogrus.Info("migrate role_api init")
	err = common.DB.AutoMigrate(RoleApi{})
	if err != nil {
		return
	}
	// feature init
	emlogrus.Info("migrate feature init")
	err = common.DB.AutoMigrate(Feature{})
	if err != nil {
		return
	}
	addAdmin := Feature{
		Name: "add",
		Menu: adminMenu.Id,
	}
	err = common.DB.FirstOrCreate(&addAdmin, addAdmin).Error
	if err != nil {
		return
	}
	deleteAdmin := Feature{
		Name: "delete",
		Menu: adminMenu.Id,
	}
	err = common.DB.FirstOrCreate(&deleteAdmin, deleteAdmin).Error
	if err != nil {
		return
	}
	lookAdmin := Feature{
		Name: "look",
		Menu: adminMenu.Id,
	}
	err = common.DB.FirstOrCreate(&lookAdmin, lookAdmin).Error
	if err != nil {
		return
	}
	// role_feature init
	emlogrus.Info("migrate role_feature init")
	err = common.DB.AutoMigrate(RoleFeature{})
	if err != nil {
		return
	}
	return
}

func Features(roles []int) (rm []Feature, err error) {
	if len(roles) == 0 {
		err = errors.New("roles not empty")
		return
	}
	root := Role{}
	err = common.DB.Where("name=?", "root").First(&root).Error
	if err != nil {
		return
	}
	if root.Id == roles[0] {
		err = common.DB.Model(Feature{}).Order("id").Find(&rm).Error
		return
	}
	for _, r := range roles {
		feas := make([]Feature, 0)
		err = common.DB.Model(Feature{}).Joins("right join role_feature on "+
			"feature.id=role_feature.feature and role_feature.role=?", r).
			Find(&feas).Error
		if err != nil {
			return
		}
		rm = append(rm, feas...)
	}
	m := make(map[int]Feature)
	fids := make([]int, 0)
	for _, fea := range rm {
		_, ok := m[fea.Id]
		if !ok {
			m[fea.Id] = fea
			fids = append(fids, fea.Id)
		}
	}
	sort.Ints(fids)
	rm = make([]Feature, 0)
	for _, fid := range fids {
		rm = append(rm, m[fid])
	}
	return
}

func UpdateRole(roleId int, mids []int, fids []int) (err error) {
	menus := make([]Menu, 0)
	common.DB.Where("role=?", roleId).Unscoped().Delete(RoleMenu{})
	common.DB.Model(Menu{}).Find(&menus)
	for _, m := range mids {
		err = common.DB.Create(&RoleMenu{Role: roleId, Menu: m}).Error
		if err != nil {
			return
		}
	}
	common.DB.Where("role=?", roleId).Unscoped().Delete(RoleFeature{})
	for _, f := range fids {
		err = common.DB.Create(&RoleFeature{Role: roleId, Feature: f}).Error
		if err != nil {
			return
		}
	}
	return
}
