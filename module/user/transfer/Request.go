package transfer

type UserRequestParam struct {
	UserId uint
}

type UserRequestBody struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginRequestBody struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
