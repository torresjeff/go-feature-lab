package featurelab

type ErrorType uint32

const (
	// Error codes returned by API, these are different from the HTTP status codes
	ErrBadRequest          ErrorType = 400
	ErrNotFound            ErrorType = 404
	ErrInternalServerError ErrorType = 500
	ErrInvalidTreatment    ErrorType = 5000
)

type Error struct {
	Code    ErrorType `json:"code"`
	Message string    `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code ErrorType, message string) *Error {
	return &Error{code, message}
}
