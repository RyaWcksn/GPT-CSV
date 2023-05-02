package dtoauthentication

type RegisterRequest struct {
	UsersParentId string `json:"parentId"`
	Name          string `json:"name" validate:"required"`
	Password      string `json:"password" validate:"required"`
	PhoneNumber   string `json:"phoneNumber" validate:"required"`
	Email         string `json:"email" validate:"required"`
	Status        int    `json:"status"`
}
