package common

// CustomError struct that implements the error interface
type CustomError struct {
	Code         int
	Message      string
	ErrorDetails string
}

// Implement the error interface by defining the Error() method
func (e *CustomError) Error() string {
	return e.ErrorDetails
}

func (e *CustomError) GetMessage() string {
	return e.Message
}

// Factory function to create a new CustomError with a default code of 400
func NewCustomError(message, details string, code ...int) *CustomError {
	errorCode := 400 // Default to 400
	if len(code) > 0 {
		errorCode = code[0]
	}

	return &CustomError{
		Code:         errorCode,
		Message:      message,
		ErrorDetails: details,
	}
}

func ValidationError(message string, code ...int) *CustomError {
	errorCode := 400 // Default to 400
	if len(code) > 0 {
		errorCode = code[0]
	}

	return &CustomError{
		Code:         errorCode,
		Message:      message,
		ErrorDetails: "",
	}
}
