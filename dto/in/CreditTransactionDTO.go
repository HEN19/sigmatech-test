package in

type CreateTransactionRequest struct {
	ContractNumber string  `json:"contract_number" binding:"required"`
	OTR            float64 `json:"otr" binding:"required,gt=0"`
	AdminFee       float64 `json:"admin_fee" binding:"required,gt=0"`
	Installment    int     `json:"installment" binding:"required,oneof=1 2 3 6"`
	Interest       float64 `json:"interest" binding:"required,gt=0"`
	AssetName      string  `json:"asset_name" binding:"required"`
	CustomerID     uint    `json:"customer_id" binding:"required"`
}
