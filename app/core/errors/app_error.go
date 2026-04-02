package errors

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Action  string `json:"action"`
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
