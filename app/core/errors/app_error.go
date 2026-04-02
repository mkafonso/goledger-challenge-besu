package errors

type AppError struct {
	Code    string
	Message string
	Action  string
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) GetAction() string {
	return e.Action
}

func (e *AppError) GetCode() string {
	return e.Code
}
