package in

import "time"

type CustomerRequest struct {
	NIK         string    `json:"NIK"`
	FullName    string    `json:"full_name"`
	LegalName   string    `json:"legal_name"`
	BirthPlace  string    `json:"birth_place"`
	BirthDate   time.Time `json:"birth_date"`
	Salary      float64   `json:"salary"`
	KTPPhoto    string    `json:"ktp_photo"`
	SelfiePhoto string    `json:"selfie_photo"`
	Limit1Month float64   `json:"limit_1_month"`
	Limit2Month float64   `json:"limit_2_month"`
	Limit3Month float64   `json:"limit_3_month"`
	Limit6Month float64   `json:"limit_6_month"`
}
