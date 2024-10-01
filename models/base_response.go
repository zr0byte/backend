package models

type BaseResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	MetaDate   interface{} `json:"metadata,omitempty"`
	Error      interface{} `json:"error,omitemoty"`
	StatusCode int         `json:"status_code"`
}
