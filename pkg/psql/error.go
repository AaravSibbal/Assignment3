package psql

type PsqlErrors string

const (
	UniqueConstraintError PsqlErrors = "23505"
	NotNullError          PsqlErrors = "23502"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessMessage struct {
	Message string `json:"message"`
}

func ConvertErrorToJsonObj(err error) *ErrorResponse {
	errReponse := ErrorResponse{
		Message: err.Error(),
	}

	return &errReponse
}
