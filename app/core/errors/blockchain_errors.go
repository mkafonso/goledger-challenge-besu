package errors

func NewErrorUnableToReadFromBlockchain() *AppError {
	return &AppError{
		Code:    "BLOCKCHAIN_READ_FAILED",
		Message: "unable to read storage from blockchain",
		Action:  "please verify blockchain connectivity and contract state",
	}
}
