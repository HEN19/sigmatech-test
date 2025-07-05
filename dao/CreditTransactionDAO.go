package dao

import (
	"database/sql"
	"fmt"

	"github.com/api-skeleton/model"
)

type creditTransactionDAO struct {
	AbstractDAO
}

var CreditTransactionDAO = creditTransactionDAO{}.New()

func (input creditTransactionDAO) New() (output creditTransactionDAO) {
	output.TableName = "transactions"
	output.FileName = "CreditTransactionDAO.go"
	return
}

func (input creditTransactionDAO) CreateTransaction(db *sql.DB, inputStruct model.CreditTransactionModel) (err error) {
	var (
		query string
	)

	query = fmt.Sprintf(`
			INSERT INTO %s
				(contract_number, otr, admin_fee, installment, interest, asset_name, customer_id)
			VALUES ($1,$2,$3,$4)
		`, input.TableName,
	)

	params := []interface{}{
		inputStruct.ContractNumber.String, inputStruct.OTR, inputStruct.AdminFee, inputStruct.Installment,
		inputStruct.Interest, inputStruct.AssetName, inputStruct.CustomerID,
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		fmt.Println(err)
		return
	}

	return

}

func (input creditTransactionDAO) GetListTransactions(db *sql.DB) (result []model.CreditTransactionModel, err error) {
	query := "SELECT id, contract_number, otr, admin_fee, installment, interest, asset_name, customer_id " +
		" FROM " + input.TableName

	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var transaction model.CreditTransactionModel
		err = rows.Scan(&transaction.ID, &transaction.ContractNumber, &transaction.OTR, &transaction.AdminFee,
			&transaction.Installment, &transaction.Interest, &transaction.AssetName, &transaction.CustomerID)
		if err != nil {
			return
		}
		result = append(result, transaction)
	}

	return
}

func (input creditTransactionDAO) GetDetailTransactions(db *sql.DB, id string) (result model.CreditTransactionModel, err error) {
	query := "SELECT id, contract_number, otr, admin_fee, installment, interest, asset_name, customer_id " +
		" FROM " + input.TableName + " WHERE id = $1"

	row := db.QueryRow(query, id)
	err = row.Scan(&result.ID, &result.ContractNumber, &result.OTR, &result.AdminFee,
		&result.Installment, &result.Interest, &result.AssetName, &result.CustomerID)
	if err != nil {
		return
	}
	return
}

func (input creditTransactionDAO) UpdateTransaction(db *sql.Tx, inputStruct model.CreditTransactionModel) (err error) {
	query := "UPDATE " + input.TableName + " SET contract_number = $1, otr = $2, admin_fee = $3, installment = $4, interest = $5, asset_name = $6, customer_id = $7 WHERE id = $8"

	params := []interface{}{
		inputStruct.ContractNumber.String, inputStruct.OTR, inputStruct.AdminFee, inputStruct.Installment,
		inputStruct.Interest, inputStruct.AssetName, inputStruct.CustomerID, inputStruct.ID,
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		return
	}
	return
}

func (input creditTransactionDAO) DeleteTransaction(db *sql.DB, id string) (err error) {
	query := "UPDATE " + input.TableName + " SET deleted = true WHERE id = $1"
	_, err = db.Exec(query, id)
	if err != nil {
		return
	}
	return
}
