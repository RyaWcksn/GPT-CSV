package dtoroles

type GetListRoleRequest struct {
	PageNumber int `json:"pageNumber" validate:"required"`
	Limit      int `json:"limit" validate:"required"`
}
