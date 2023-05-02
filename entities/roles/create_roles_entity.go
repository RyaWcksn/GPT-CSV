package entityroles

type CreateRoleDetails struct {
	Topic            string `json:"topic"`
	ChildDescription string `json:"childDescription"`
	RoleName         string `json:"roleName"`
	RoleDescription  string `json:"roleDescription"`
}
