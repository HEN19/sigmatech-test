package in

import (
	"time"

	"github.com/api-skeleton/constanta/ErrorModel"
	"github.com/gin-gonic/gin"
)

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

// before inserting a new customer. It checks for mandatory fields and returns
func (input *CustomerRequest) ValidationCustomerInsert(c *gin.Context) ErrorModel.DynamicErrorResponse {
	return input.mandatoryValidation()
}

func (input *CustomerRequest) mandatoryValidation() ErrorModel.DynamicErrorResponse {
	if input.NIK == "" {
		return ErrorModel.ErrorInvalidRequest(nil, "NIK", "NIK is required")
	}
	if input.FullName == "" {
		ErrorModel.ErrorInvalidRequest(nil, "Fullname", "Fullname is required")
	}

	return ErrorModel.NonErrorResponse()
}
