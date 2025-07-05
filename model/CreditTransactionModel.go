package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type CreditTransactionModel struct {
	ID             uuid.UUID
	ContractNumber sql.NullString  // e.g. 'CN12345'
	OTR            sql.NullFloat64 // e.g. On-the-road price
	AdminFee       sql.NullFloat64
	Installment    sql.NullInt64 // e.g. number of months
	Interest       sql.NullFloat64
	AssetName      sql.NullString
	CustomerID     uuid.UUID // change to nullable if necessary
	Customer       CustomerModel
}