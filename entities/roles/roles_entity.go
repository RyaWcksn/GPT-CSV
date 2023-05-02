package entityroles

type RoleDetails struct {
	Topic            string `json:"topic"`
	Rules            string `json:"rules"`
	Goals            string `json:"goals"`
	ChildDescription string `json:"childDescription"`
	RoleName         string `json:"roleName"`
	RoleDescription  string `json:"roleDescription"`
}
