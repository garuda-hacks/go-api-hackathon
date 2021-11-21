package response

type InterestMap struct {
	ID         int        `json:"id"`
	IDInterest int        `json:"id_interest"`
	IDUser     int        `json:"id_user"`
	Interest   []Interest `json:"interest"`
}
