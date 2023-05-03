package dtochild

type CreateUserChildRequest struct {
	ParentId  string `json:"parentId"`
	ChildName string `json:"childName" validate:"required"`
	RoleName  string `json:"roleName"`
	Age       int    `json:"age" validate:"required"`
}
