package types

// struct tag meaning
// `marshaling struct type: name, [omitempty]`
type ApiError struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

//
type ApiResponse struct {
	Result interface{} `json:"result"`
	Error ApiError `json:"error"`
	TraceId string `json:"traceId"`
}
