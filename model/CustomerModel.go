package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type CustomerModel struct {
	ID          uuid.UUID      // UUID is non-nullable, generated manually
	NIK         sql.NullString // Assume not nullable and unique
	FullName    sql.NullString
	LegalName   sql.NullString
	BirthPlace  sql.NullString
	BirthDate   sql.NullString // Or use sql.NullTime if date parsing is needed
	Salary      sql.NullFloat64
	KTPPhoto    sql.NullString
	SelfiePhoto sql.NullString
	Limit1Month sql.NullFloat64
	Limit2Month sql.NullFloat64
	Limit3Month sql.NullFloat64
	Limit6Month sql.NullFloat64
}
