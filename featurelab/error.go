package featurelab

type ErrorType uint32

const (
	ErrBadRequest          ErrorType = 400
	ErrNotFound            ErrorType = 404
	ErrInternalServerError ErrorType = 500
)

type Error struct {
	Code    ErrorType `json:"code"`
	Message string    `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code ErrorType, message string) Error {
	return Error{code, message}
}
