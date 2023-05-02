package dtoroles

type CreateRoleRequest struct {
	RoleId           string `json:"roleId"`
	Topic            string `json:"topic" validate:"required"`
	Rules            string `json:"rules" validate:"required"`
	Goals            string `json:"goals" validate:"required"`
	ChildDescription string `json:"childDescription" validate:"required"`
	RoleName         string `json:"roleName" validate:"required"`
	RoleDescription  string `json:"roleDescription" validate:"required"`
}
