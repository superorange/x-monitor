package api

type Response struct {
	Code    int   `json:"code"`
	Data    any   `json:"data,omitempty"`
	Message any   `json:"message,omitempty"`
	Error   error `json:"error,omitempty"`
}

func NewApiResponse(code int, data any) *Response {
	return &Response{
		Code: code,
		Data: data,
	}
}
