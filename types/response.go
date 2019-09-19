package types

type Error struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Info    interface{} `json:"info,omitempty"`
}

type Response struct {
	Success bool        `json:"success"`
	Error   interface{} `json:"error,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}
