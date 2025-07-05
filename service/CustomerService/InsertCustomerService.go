package CustomerService

import (
	"database/sql"

	"github.com/api-skeleton/config"
	"github.com/api-skeleton/constanta/ErrorModel"
	"github.com/api-skeleton/dao"
	"github.com/api-skeleton/dto/in"
	"github.com/api-skeleton/model"
	"github.com/google/uuid"
)

func InsertCustomerService(req *in.CustomerRequest) (*string, ErrorModel.DynamicErrorResponse) {
	var reqBody *model.CustomerModel
	// Map the request body to the model
	reqBody = mapToCustomerModel(req)

	// Connect to DB (use connection pool for better performance)
	db, errTx := config.Connect().Begin()
	if errTx != nil {
		db.Rollback()
		return nil, ErrorModel.ErrorInternalServerError(nil, errTx.Error())
	}
	defer db.Commit()

	// Insert user into DB
	tx := dao.CustomerDAO.CreateCustomer(db, reqBody)
	if tx != nil {
		return nil, ErrorModel.ErrorInternalServerError(nil, tx.Error())
	}

	// Respond with success
	return nil, ErrorModel.NonErrorResponse()
}

func mapToCustomerModel(req *in.CustomerRequest) *model.CustomerModel {
	return &model.CustomerModel{
		ID:          uuid.New(),
		NIK:         sql.NullString{String: req.NIK},
		FullName:    sql.NullString{String: req.FullName},
		LegalName:   sql.NullString{String: req.LegalName},
		BirthPlace:  sql.NullString{String: req.BirthPlace},
		BirthDate:   sql.NullString{String: req.BirthDate.Format("2006-01-02")},
		Salary:      sql.NullFloat64{Float64: req.Salary},
		KTPPhoto:    sql.NullString{String: req.KTPPhoto},
		SelfiePhoto: sql.NullString{String: req.SelfiePhoto},
		Limit1Month: sql.NullFloat64{Float64: req.Limit1Month},
		Limit2Month: sql.NullFloat64{Float64: req.Limit2Month},
		Limit3Month: sql.NullFloat64{Float64: req.Limit3Month},
		Limit6Month: sql.NullFloat64{Float64: req.Limit6Month},
	}
}
