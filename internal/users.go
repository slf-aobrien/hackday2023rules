package rules

type Users struct {
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Birthday  string `json:"birthday"`
	Gender    string `json:"gender"`
	OcCode    string `json:"ocCode"`
	HasDental bool   `json:"HasDental"`
	HasLife   bool   `json:"HasLife"`
	HasLtd    bool   `json:"HasLtd"`
}
