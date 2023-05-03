package dtoroles

type UpdateSingleRoleRequest struct {
	ParentId         string `json:"parentId"`
	RoleName         string `json:"roleName" validate:"required"`
	Topic            string `json:"topic" validate:"required"`
	Rules            string `json:"rules" validate:"required"`
	Goals            string `json:"goals" validate:"required"`
	ChildDescription string `json:"childDescription" validate:"required"`
	RoleDescription  string `json:"roleDescription" validate:"required"`
}
