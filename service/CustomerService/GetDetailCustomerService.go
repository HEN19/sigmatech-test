package CustomerService

import (
	"time"

	"github.com/api-skeleton/config"
	"github.com/api-skeleton/constanta/ErrorModel"
	"github.com/api-skeleton/dao"
	"github.com/api-skeleton/dto/out"
)

func GetDetailCustomer(id string) (*out.CustomerDTOOut, ErrorModel.DynamicErrorResponse) {
	var responses out.CustomerDTOOut
	db := config.Connect()
	defer db.Close()

	// Get user profile from the DAO layer using the decoded user ID
	customers, errGet := dao.CustomerDAO.GetDetailCustomer(db, id)
	if errGet != nil {
		return nil, ErrorModel.ErrorDataNotFound(errGet.Error())
	}

	birthDate, _ := time.Parse("2006-01-02", customers.BirthDate.String)
	responses = out.CustomerDTOOut{
		ID:          customers.ID,
		NIK:         customers.NIK.String,
		FullName:    customers.FullName.String,
		LegalName:   customers.LegalName.String,
		BirthPlace:  customers.BirthPlace.String,
		BirthDate:   birthDate,
		Salary:      customers.Salary.Float64,
		KTPPhoto:    customers.KTPPhoto.String,
		SelfiePhoto: customers.SelfiePhoto.String,
		Limit1Month: customers.Limit1Month.Float64,
		Limit2Month: customers.Limit2Month.Float64,
		Limit3Month: customers.Limit3Month.Float64,
		Limit6Month: customers.Limit6Month.Float64,
	}
	return &responses, ErrorModel.NonErrorResponse()
}
