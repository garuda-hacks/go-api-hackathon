package response

type Ugcprofile struct {
	IDUser    int    `json:"id_user"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}
