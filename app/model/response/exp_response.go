package response

type Experiences struct {
	ID              int    `json:"id"`
	IDUser          int    `json:"id_user"`
	Title           string `json:"title"`
	EmployementType string `json:"employementtype"`
	CompanyName     string `json:"companyname"`
	City            string `json:"city"`
	Country         string `json:"country"`
	StartDate       string `json:"startdate"`
	EndDate         string `json:"enddate"`
	Description     string `json:"description"`
}
