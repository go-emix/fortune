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
