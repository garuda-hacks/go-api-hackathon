package response

type Internship struct {
	ID          int    `json:"id"`
	IDIntern    int    `json:"id_intern"`
	Title       string `json:"judul" copier:"Title"`
	Provider    string `json:"provider"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Description string `json:"description"`
	Requirement string `json:"requirement"`
	StartDate   string `json:"startdate" copier:"StartDate"`
	EndDate     string `json:"enddate" copier:"EndDate"`
	IsDone      string `json:"isdone"`
	Avatar      string `json:"avatar"`
	Target      string `json:"target" copier:"Target"`
	Duration    int    `json:"duration" copier:"Duration"`
}
