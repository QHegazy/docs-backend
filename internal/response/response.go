package response

type BaseResponse struct {
	Status  int    `int:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	BaseResponse
	Data map[string]interface{} `json:"data"`
}

type ErrorResponse struct {
	BaseResponse
	Error any `json:"error"`
}
