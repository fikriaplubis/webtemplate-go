package dto

type DirectTransferDetailTempRequest struct {
	NoReff       string  `json:"no_reff"`
	Sequence     string  `json:"sequence"`
	CIF          string  `json:"cif"`
	EmployeeName string  `json:"employee_name"`
	Currency     string  `json:"currency"`
	AccountNo    string  `json:"account_no"`
	Amount       float64 `json:"amount"`
	Description  string  `json:"description"`
}
