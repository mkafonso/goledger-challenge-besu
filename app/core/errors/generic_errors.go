package errors

func NewInternalError() *AppError {
	return &AppError{
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "internal server error",
		Action:  "please try again later or contact support",
	}
}
