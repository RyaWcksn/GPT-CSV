package entitychat

type CreateNewChatDetail struct {
	ChildName string `json:"childName"`
	RoleName  string `json:"roleName"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
}
