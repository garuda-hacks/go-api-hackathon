package response

type Scholarship struct {
	ID            int    `json:"id"`
	IDScholarship int    `json:"id_scholarship"`
	Title         string `json:"judul" copier:"Title"`
	Provider      string `json:"provider"`
	Location      string `json:"location"`
	Description   string `json:"description"`
	Requirement   string `json:"requirement"`
	StartDate     string `json:"startdate" copier:"StartDate"`
	EndDate       string `json:"enddate" copier:"EndDate"`
	IsDone        bool   `json:"isdone"`
	Avatar        string `json:"avatar"`
	Target        string `json:"target" copier:"Target"`
	Duration      int    `json:"duration" copier:"Duration"`
}
