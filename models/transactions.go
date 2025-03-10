package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID               uuid.UUID `json:"id"`
	UserID           uuid.UUID `json:"user_id"`
	Tenor            int       `json:"tenor"`
	ContractNumber   string    `json:"contract_number"`
	OTR              float64   `json:"otr"`
	AdminFee         float64   `json:"admin_fee"`
	InstallmentCount int       `json:"installment_count"`
	Interest         float64   `json:"interest"`
	AssetName        string    `json:"asset_name"`
	AssetType        string    `json:"asset_type"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
