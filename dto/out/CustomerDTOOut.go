package out

import (
	"time"

	"github.com/google/uuid"
)

type CustomerDTOOut struct {
	ID          uuid.UUID
	NIK         string
	FullName    string
	LegalName   string
	BirthPlace  time.Time
	BirthDate   time.Time
	Salary      float64
	KTPPhoto    string
	SelfiePhoto string
	Limit1Month float64
	Limit2Month float64
	Limit3Month float64
	Limit6Month float64
}
