package response

type BaseResponse struct {
	Status  int    `json:"status"` // Corrected tag
	Message string `json:"message"`
}

type SuccessResponse struct {
	BaseResponse
	Data any `json:"data"`
}

type ErrorResponse struct {
	BaseResponse
	Error any `json:"error"`
}
