package response

type BookmarkScholarship struct {
	ID            int    `json:"id"`
	IDScholarship int    `json:"id_scholarship"`
	Title         string `json:"judul" copier:"Title"`
	EndDate       string `json:"enddate" copier:"EndDate"`
	Avatar        string `json:"avatar"`
	Target        string `json:"target" copier:"Target"`
	Duration      int    `json:"duration" copier:"Duration"`
}
