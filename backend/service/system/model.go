package system

type LoginParam struct {
	Username string
	Password string
}

type User struct {
	Id       int
	Username string
	Password string
	Nickname string
	Roles    []Role `gorm:"-"`
}

func (User) TableName() string {
	return "user"
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
	Auth      string `json:"auth"`
	Parent    int    `json:"parent"`
}

func (Menu) TableName() string {
	return "menu"
}

type Role struct {
	Id    int
	Name  string
	Menus []Menu `gorm:"-"`
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

type UserRole struct {
	Id   int
	User int
	Role int
}

func (UserRole) TableName() string {
	return "user_role"
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
