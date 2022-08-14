package system

type LoginParam struct {
	Username string
	Password string
}

// Admin 管理员
type Admin struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Roles    []Role `gorm:"-" json:"roles"`
	// 是否启用
	Enabled bool `json:"enabled"`
}

func (Admin) TableName() string {
	return "admin"
}

func (r Admin) ToSimple() SimpleAdmin {
	var sa = SimpleAdmin{
		Id:       r.Id,
		Username: r.Username,
		Nickname: r.Nickname,
		Enabled:  r.Enabled,
		Rids:     make([]int, 0),
	}
	for _, r := range r.Roles {
		sa.Rids = append(sa.Rids, r.Id)
	}
	return sa
}

type SimpleAdmin struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Rids     []int  ` json:"rids"`
	Enabled  bool   `json:"enabled"`
}

type LoginSucc struct {
	Id       int    `json:"id"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

type Menu struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Component string `json:"component"`
	Path      string `json:"path"`
	// 权限标识
	Auth   string `json:"auth"`
	Parent int    `json:"parent"`
}

func (Menu) TableName() string {
	return "menu"
}

type Role struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Menus []Menu `gorm:"-" json:"-"`
}

func (Role) TableName() string {
	return "role"
}

type RoleMenu struct {
	Id   int
	Role int
	Menu int
}

func (RoleMenu) TableName() string {
	return "role_menu"
}

type AdminRole struct {
	Id    int
	Admin int
	Role  int
}

func (AdminRole) TableName() string {
	return "admin_role"
}

type Api struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

func (Api) TableName() string {
	return "api"
}

type RoleApi struct {
	Id   int
	Role int
	Api  int
}

func (RoleApi) TableName() string {
	return "role_api"
}

type Feature struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Menu       int    `json:"menu_id"`
	MenuEntity Menu   `json:"menu" gorm:"-"`
}

func (Feature) TableName() string {
	return "feature"
}

type RoleFeature struct {
	Id      int
	Role    int
	Feature int
}

func (RoleFeature) TableName() string {
	return "role_feature"
}

// FrontI18N 前端i18n
type FrontI18N struct {
	Name string
	En   string
	Zh   string
}

func (FrontI18N) TableName() string {
	return "front_i18n"
}
