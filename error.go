package citcallgo

// citcallgo error code type
type ErrorCode string

// citcallgo error code struct
type Error struct {
	Status  int       `json:"status"`
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}
