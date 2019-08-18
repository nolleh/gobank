package types

// struct tag meaning
// `marshaling struct type: name,[omitempty]` (no space)
type ApiError struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

//
type ApiResponse struct {
	Result interface{} `json:"result,omitempty"`
	Error *ApiError `json:"error,omitempty"`
	TraceId string `json:"traceId"`
}
