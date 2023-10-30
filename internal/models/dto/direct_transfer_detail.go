package dto

type DirectTrasferDetailRequest struct {
	NoReff       string  `json:"no_reff"`
	NoUrut       string  `json:"no_urut"`
	EmployeeName string  `json:"employee_name"`
	Currency     string  `json:"currency"`
	AccountNo    string  `json:"account_no"`
	Amount       float64 `json:"amount"`
	Description  string  `json:"description"`
	CIF          string  `json:"cif"`
}
