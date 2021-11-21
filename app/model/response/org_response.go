package response

type Organization struct {
	ID          int    `json:"id"`
	IDUser      int    `json:"id_user"`
	Institute   string `json:"institute"`
	Location    string `json:"location"`
	Role        string `json:"role"`
	Description string `json:"description"`
	StartDate   string `json:"startdate"`
	EndDate     string `json:"enddate"`
	Photo       string `json:"photo"`
}
