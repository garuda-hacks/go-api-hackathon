package response

type UgcList struct {
	ID           int          `json:"id" copier:"ID"`
	IDUgc        int          `json:"id_ugc" copier:"IDUgc"`
	IDUser       int          `json:"id_user" copier:"IDUser"`
	IDInterest   int          `json:"id_interest" copier:"IDInterest"`
	Posting      string       `json:"posting" copier:"Posting"`
	PostingDate  string       `json:"postingdate" copier:"PostingDate"`
	PhotoPosting string       `json:"photopost" copier:"PhotoPosting"`
	VideoPosting string       `json:"videopost" copier:"VideoPosting"`
	Ugcprofile   []Ugcprofile `json:"profile"`
}
