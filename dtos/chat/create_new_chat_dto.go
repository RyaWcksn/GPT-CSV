package dtochat

type CreateNewChatRequest struct {
	ParentId  string `json:"parentId"`
	ChildName string `json:"childName" validate:"required"`
	RoleName  string `json:"roleName" validate:"required"`
	Question  string `json:"question" validate:"required"`
	Answer    string `json:"answer"`
}
