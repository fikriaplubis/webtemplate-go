package custom

type CustomError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}
