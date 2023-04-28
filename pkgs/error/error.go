package customerror

import "fmt"

// List of all errors
const (
	InternalServer     = "internal error"
	BadRequest         = "bad request"
	UnavailableService = "unavailable service"
	Unauthorized       = "unauthorized"
	RequestNotFound    = "not found"
)

const (
	UnavailableServiceCode = 503
	InternalServerCode     = 500
	BadRequestCode         = 400
	UnauthorizedCode       = 401
	ReuestNotFoundCode     = 404
)

type IError interface {
	Error() string
	GetHTTPCode() int
}

type ErrorForm struct {
	Code            int                     `json:"code"`
	Message         string                  `json:"message"`
	CommonError     string                  `json:"error,omitempty"`
	ValidationError []ErrorValidatorDetails `json:"errorValidate,omitempty"`
}

type ErrorValidatorDetails struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (o ErrorForm) Error() string {
	if o.CommonError == "" && o.ValidationError != nil {
		return fmt.Sprintf("CustomError code = %v desc - %v errors = %v", o.Code, o.Message, o.ValidationError)
	}
	return fmt.Sprintf("CustomError code = %v desc - %v errors = %v", o.Code, o.Message, o.CommonError)
}

func (o ErrorForm) GetHTTPCode() int {
	return o.Code
}

// getErrorCode return http error code base on error message.
func getErrorCode(errMsg string) int {
	var val int
	switch errMsg {
	case InternalServer:
		val = InternalServerCode
	case UnavailableService:
		val = UnavailableServiceCode
	case BadRequest:
		val = BadRequestCode
	case Unauthorized:
		val = UnauthorizedCode
	case RequestNotFound:
		val = ReuestNotFoundCode
	default:
		val = InternalServerCode
	}
	return val
}

// GetError code and message then return.
func GetError(errMessage string, errActual error) *ErrorForm {
	return &ErrorForm{
		Code:        getErrorCode(errMessage),
		Message:     errMessage,
		CommonError: errActual.Error(),
	}
}

// GetErrorValidation code and message then return.
func GetErrorValidation(errMessage string, errs []ErrorValidatorDetails) *ErrorForm {
	return &ErrorForm{
		Code:            getErrorCode(errMessage),
		Message:         errMessage,
		ValidationError: errs,
	}
}
