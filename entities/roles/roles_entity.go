package entityroles

type RoleDetails struct {
	RoleName         string `json:"roleName"`
	Topic            string `json:"topic"`
	Rules            string `json:"rules"`
	Goals            string `json:"goals"`
	ChildDescription string `json:"childDescription"`
	RoleDescription  string `json:"roleDescription"`
}
