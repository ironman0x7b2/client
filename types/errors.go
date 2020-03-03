package types

type Error struct {
	Module  string `json:"module,omitempty"`
	Code    uint32 `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewError(module string, code uint32, message string) *Error {
	return &Error{
		Module:  module,
		Code:    code,
		Message: message,
	}
}
