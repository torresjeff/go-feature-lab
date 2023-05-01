package featurelab

type ErrorType uint32

const (
	// Error codes returned by API, these are different from the HTTP status codes
	ErrBadRequest          uint32 = 400
	ErrNotFound            uint32 = 404
	ErrInternalServerError uint32 = 500
	ErrInvalidTreatment    uint32 = 5000
)

type Error struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code uint32, message string) *Error {
	return &Error{code, message}
}
