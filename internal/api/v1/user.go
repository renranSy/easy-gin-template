package v1

type LoginReq struct {
	Username string `form:"username" bind:"required"`
	Password string `form:"password" bind:"required"`
}

type LoginRes struct {
	UserInfo UserInfo `json:"userInfo"`
	Token    string   `json:"token"`
}

type UserInfo struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}
