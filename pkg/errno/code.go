package errno

var (
	// Common errors
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	// user errors
	ErrUserNotFound = &Errno{Code: 20102, Message: "The user was not found."}
	ErrDatabase     = &Errno{Code: 20111, Message: "ssss"}
)
