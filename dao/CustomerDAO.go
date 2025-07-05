package dao

import (
	"database/sql"
	"fmt"

	"github.com/api-skeleton/model"
)

type customerDAO struct {
	AbstractDAO
}

var CustomerDAO = customerDAO{}.New()

func (input customerDAO) New() (output customerDAO) {
	output.TableName = "customers"
	output.FileName = "CustomerDAO.go"
	return
}

func (input customerDAO) CreateCustomer(db *sql.Tx, inputStruct *model.CustomerModel) (err error) {
	var (
		query string
	)

	query = fmt.Sprintf(`
			INSERT INTO %s
				(NIK, full_name, legal_name, 
				birth_place, birth_date, salary, 
				ktp_photo,selfie_photo,
				limit_1_month,limit_2_month,limit_3_month,limit_6_month)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
		`, input.TableName,
	)

	params := []interface{}{
		inputStruct.NIK.String, inputStruct.FullName.String, inputStruct.LegalName.String,
		inputStruct.BirthPlace.String, inputStruct.BirthDate.String, inputStruct.Salary.Float64,
		inputStruct.KTPPhoto.String, inputStruct.SelfiePhoto.String,
		inputStruct.Limit1Month.Float64, inputStruct.Limit2Month.Float64, inputStruct.Limit3Month.Float64, inputStruct.Limit6Month.Float64,
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		return
	}

	return
}

func (input customerDAO) GetListCustomer(db *sql.DB, page, limit int) (result *[]model.CustomerModel, err error) {
	var (
		customer  model.CustomerModel
		customers []model.CustomerModel
	)
	query := "SELECT id, NIK, full_name, legal_name, birth_place, birth_date, salary, ktp_photo,selfie_photo,limit_1_month,limit_2_month,limit_3_month,limit_6_month FROM " + input.TableName + "LIMIT ? OFFSET ?`"

	rows, err := db.Query(query, limit, page)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&customer.ID, &customer.NIK, &customer.FullName, &customer.LegalName,
			&customer.BirthPlace, &customer.BirthDate, &customer.Salary, &customer.KTPPhoto, &customer.SelfiePhoto,
			&customer.Limit1Month, &customer.Limit2Month, &customer.Limit3Month, &customer.Limit6Month)
		if err != nil {
			return
		}
		customers = append(customers, customer)
	}

	return &customers, nil
}

func (input customerDAO) GetDetailCustomer(db *sql.DB, id string) (result model.CustomerModel, err error) {
	query := "SELECT id, NIK, full_name, legal_name, birth_place, birth_date, salary, ktp_photo,selfie_photo,limit_1_month,limit_2_month,limit_3_month,limit_6_month FROM " + input.TableName + " WHERE id = $1"

	row := db.QueryRow(query, id)
	err = row.Scan(&result.ID, &result.NIK, &result.FullName, &result.LegalName,
		&result.BirthPlace, &result.BirthDate, &result.Salary, &result.KTPPhoto, &result.SelfiePhoto,
		&result.Limit1Month, &result.Limit2Month, &result.Limit3Month, &result.Limit6Month)
	if err != nil {
		return
	}
	return
}

func (input customerDAO) UpdateCustomer(db *sql.Tx, inputStruct model.CustomerModel) (err error) {
	query := "UPDATE " + input.TableName + " SET NIK = $1, full_name = $2, legal_name = $3, " +
		"birth_place = $4, birtd_date = $5, salary = $6, ktp_photo = $7, selfie_photo = $8, " +
		"limit_1_month = $9,limit_2_month = $10,limit_3_month = $11,limit_6_month = $12 WHERE id = $13"

	params := []interface{}{
		inputStruct.NIK.String, inputStruct.FullName.String, inputStruct.LegalName.String,
		inputStruct.BirthPlace.String, inputStruct.BirthDate.String, inputStruct.Salary,
		inputStruct.KTPPhoto.String, inputStruct.SelfiePhoto.String,
		inputStruct.Limit1Month, inputStruct.Limit2Month, inputStruct.Limit3Month, inputStruct.Limit6Month,
		inputStruct.ID,
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		return
	}

	return
}

func (input customerDAO) DeleteCustomer(db *sql.DB, id string) (err error) {
	query := "UPDATE " + input.TableName + " SET deleted = true WHERE id = $1"
	_, err = db.Exec(query, id)
	if err != nil {
		return
	}
	return
}
