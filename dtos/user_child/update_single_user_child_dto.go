package dtochild

type UpdateSingleUserChildRequest struct {
	ParentId  string `json:"parentId"`
	ChildName string `json:"childName"`
	RoleName  string `json:"roleName"`
	Age       int    `json:"age"`
}
