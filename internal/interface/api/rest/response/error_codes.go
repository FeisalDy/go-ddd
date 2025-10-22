package response

const (
	// General errors
	CodeInternalServerError = "GEN001"
	CodeBadRequest          = "GEN002"
	CodeNotFound            = "GEN003"
	CodeValidationError     = "GEN004"

	// User domain
	CodeUserAlreadyExists = "USR001"
	CodeUserNotFound      = "USR002"

	// Auth domain
	CodeInvalidCredentials = "AUTH001"
	CodeTokenExpired       = "AUTH002"
	CodeUnauthorized       = "AUTH003"
)
