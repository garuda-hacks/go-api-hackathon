package response

type BookmarkInternship struct {
	ID       int    `json:"id"`
	IDIntern int    `json:"id_intern"`
	Title    string `json:"judul" copier:"Title"`
	EndDate  string `json:"enddate" copier:"EndDate"`
	Avatar   string `json:"avatar"`
	Target   string `json:"target" copier:"Target"`
	Duration int    `json:"duration" copier:"Duration"`
}
