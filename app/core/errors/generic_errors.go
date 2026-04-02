package errors

func NewInternalError() *AppError {
	return &AppError{
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "internal server error",
		Action:  "please try again later or contact support",
	}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Code:    "BAD_REQUEST",
		Message: message,
		Action:  "please check request payload and try again",
	}
}
