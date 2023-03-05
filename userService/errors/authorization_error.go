package errors

type authorization_error struct {
	Error *Base_error `json:"error"`
}

func New_authorization_error(message string, err error) *authorization_error {
	p := new(authorization_error)
	p.Error = New_Base_error(AUTHORIZATION_ERROR, 403, message, err)
	return p
}