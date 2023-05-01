package featurelab

type ErrorType uint8

const (
	ErrBadRequest          = 400
	ErrNotFound            = 404
	ErrInternalServerError = 500
)

type Error struct {
	code    ErrorType
	message string
}

func (e *Error) Code() ErrorType {
	return e.code
}

func (e *Error) Error() string {
	return e.message
}

func NewError(code ErrorType, message string) Error {
	return Error{code, message}
}
