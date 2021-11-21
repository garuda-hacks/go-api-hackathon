package response

type Education struct {
	ID           int    `json:"id"`
	IDUser       int    `json:"id_user"`
	School       string `json:"school"`
	Degree       string `json:"degree"`
	FieldofStudy string `json:"fieldofstudy"`
	StartDate    string `json:"startdate"`
	EndDate      string `json:"enddate"`
	Description  string `json:"description"`
	Photo        string `json:"photo"`
}
