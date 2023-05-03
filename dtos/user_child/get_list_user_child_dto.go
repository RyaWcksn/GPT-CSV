package dtochild

type GetListUserChildRequest struct {
	PageNumber int `json:"pageNumber" validate:"required"`
	Limit      int `json:"limit" validate:"required"`
}
