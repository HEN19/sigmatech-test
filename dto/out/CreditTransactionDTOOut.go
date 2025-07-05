package out

import "github.com/google/uuid"

type TransactionResponse struct {
	ID             uuid.UUID `json:"id"`
	ContractNumber string    `json:"contract_number"`
	OTR            float64   `json:"otr"`
	AdminFee       float64   `json:"admin_fee"`
	Installment    int       `json:"installment"`
	Interest       float64   `json:"interest"`
	AssetName      string    `json:"asset_name"`
	CustomerID     uint      `json:"customer_id"`
}
