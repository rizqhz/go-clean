package transfer

type UserResponseBody struct {
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token,omitempty"`
}

type LoginResponseBody struct {
	Token string `json:"token,omitempty"`
}
