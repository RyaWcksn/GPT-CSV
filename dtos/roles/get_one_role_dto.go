package dtoroles

type GetOneRoleRequest struct {
	ParentId string `json:"parentId"`
	RoleName string `json:"roleName" validate:"required"`
}
