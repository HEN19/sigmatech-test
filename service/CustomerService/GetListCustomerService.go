package CustomerService

import (
	"time"

	"github.com/api-skeleton/config"
	"github.com/api-skeleton/constanta/ErrorModel"
	"github.com/api-skeleton/dao"
	"github.com/api-skeleton/dto/out"
)

func GetListCustomerService(page, limit int) (*[]out.CustomerDTOOut, ErrorModel.DynamicErrorResponse) {
	// Connect to DB (use connection pool for better performance)
	var responses []out.CustomerDTOOut
	db := config.Connect()
	defer db.Close()

	offset := (page - 1) * limit

	// Get user profile from the DAO layer using the decoded user ID
	customers, errGet := dao.CustomerDAO.GetListCustomer(db, offset, limit)
	if errGet != nil {
		return nil, ErrorModel.ErrorDataNotFound(errGet.Error())
	}

	for _, customer := range *customers {
		birthdate, _ := time.Parse("2006-01-02", customer.BirthDate.String)
		resp := out.CustomerDTOOut{
			ID:          customer.ID,
			NIK:         customer.NIK.String,
			FullName:    customer.FullName.String,
			LegalName:   customer.LegalName.String,
			BirthPlace:  customer.BirthPlace.String,
			BirthDate:   birthdate,
			Salary:      customer.Salary.Float64,
			KTPPhoto:    customer.KTPPhoto.String,
			SelfiePhoto: customer.SelfiePhoto.String,
			Limit1Month: customer.Limit1Month.Float64,
			Limit2Month: customer.Limit2Month.Float64,
			Limit3Month: customer.Limit3Month.Float64,
			Limit6Month: customer.Limit6Month.Float64,
		}

		responses = append(responses, resp)
	}
	return &responses, ErrorModel.NonErrorResponse()
}
