package dtos

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
