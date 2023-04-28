package entityauthentication

type RegisterDetails struct {
	UsersParentId string `json:"parentId"`
	Name          string `json:"name"`
	Password      string `json:"-"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phoneNumber"`
	Status        int    `json:"status"`
}
