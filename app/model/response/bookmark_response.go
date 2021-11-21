package response

type Bookmark struct {
	ID            int                   `json:"id" copier:"ID"`
	IDBookmark    int                   `json:"id_bookmark" copier:"IDBookmark"`
	IDUser        int                   `json:"id_user" copier:"IDUser"`
	IDIntern      int                   `json:"id_intern" copier:"IDIntern"`
	IDScholarship int                   `json:"id_scholarship" copier:"IDScholarship"`
	Name          string                `json:"name" copier:"Name"`
	Internship    []BookmarkInternship  `json:"internship"`
	Scholarship   []BookmarkScholarship `json:"scholarship"`
}
