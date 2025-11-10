package psql


// unfortunate reality of enums in go
type PsqlErrors string
const (
	UniqueConstraintError PsqlErrors = "23505"
	NotNullError          PsqlErrors = "23502"
)

// for error response and its json things
type ErrorResponse struct {
	Message string `json:"message"`
}

// for success response and its json things
type SuccessMessage struct {
	Message string `json:"message"`
}

// for converting error Obj to json (costom)
func ConvertErrorToJsonObj(err error) *ErrorResponse {
	errReponse := ErrorResponse{
		Message: err.Error(),
	}

	return &errReponse
}
