package response

type Ugc struct {
	ID           int    `json:"id"`
	IDUgc        int    `json:"id_ugc"`
	IDUser       int    `json:"id_user"`
	IDInterest   int    `json:"id_interest"`
	Posting      string `json:"posting"`
	PostingDate  string `json:"postingdate"`
	PhotoPosting string `json:"photopost"`
	VideoPosting string `json:"videopost"`
}
