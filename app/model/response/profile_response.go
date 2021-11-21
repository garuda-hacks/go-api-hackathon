package response

type ResultProfil struct {
	ID           int            `json:"id"`
	IDUser       int            `json:"id_user"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	Handphone    string         `json:"handphone"`
	Avatar       string         `json:"avatar"`
	Following    int            `json:"following"`
	Followers    int            `json:"followers"`
	Community    int            `json:"community"`
	Instagram    string         `json:"instagram"`
	Facebook     string         `json:"facebook"`
	Twitter      string         `json:"twitter"`
	Linkedin     string         `json:"linkedin"`
	Experiences  []Experiences  `json:"experiences"`
	Education    []Education    `json:"education"`
	Organization []Organization `json:"organization"`
	UGC          []Ugc          `json:"ugc"`
	Interest     []InterestMap  `json:"interest_map"`
}
type JsonResultProfil struct {
	User ResultProfil `json:"user"`
}
