package system

type LoginParam struct {
	Username string
	Password string
}

// Admin 管理员
type Admin struct {
	Id       int
	Username string
	Password string
	Nickname string
	Roles    []Role `gorm:"-"`
	// 是否启用
	Enabled bool
}

func (Admin) TableName() string {
	return "admin"
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
	Id   int
	Name string
	Path string
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
	Id   int    `json:"id"`
	Name string `json:"name"`
	Menu int    `json:"menu"`
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
