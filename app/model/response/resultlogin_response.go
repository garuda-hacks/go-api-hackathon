package response

type ResultUser struct {
	ID       int    `json:"id"`
	IDUser   int    `json:"id_user"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ResponeRegister struct {
	User       interface{} `json:"user"`
	Token      string      `json:"token"`
	JwtToken   string      `json:"jwt_token"`
	ServerTime string      `json:"server_time"`
}
