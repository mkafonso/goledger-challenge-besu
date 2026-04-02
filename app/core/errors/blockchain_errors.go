package errors

func NewErrorUnableToReadFromBlockchain() *AppError {
	return &AppError{
		Code:    "BLOCKCHAIN_READ_FAILED",
		Message: "unable to read storage from blockchain",
		Action:  "please verify blockchain connectivity and contract state",
	}
}

func NewErrorUnableToWriteToBlockchain() *AppError {
	return &AppError{
		Code:    "BLOCKCHAIN_WRITE_FAILED",
		Message: "unable to write storage to blockchain",
		Action:  "please verify transaction parameters and try again",
	}
}
