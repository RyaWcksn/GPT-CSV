package entityauthentication

type LoginDetails struct {
	ParentId   string `json:"parentId"`
	Token      string `json:"loginToken"`
	ExpiryDate int64  `json:"expiryDate"`
}
