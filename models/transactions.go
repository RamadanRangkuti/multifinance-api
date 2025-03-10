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

type TransactionRequest struct {
	UserID           uuid.UUID `json:"user_id"`
	Tenor            int       `json:"tenor" validate:"required"`
	ContractNumber   string    `json:"contract_number,omitempty"`
	OTR              float64   `json:"otr" validate:"required"`
	AdminFee         float64   `json:"admin_fee" validate:"required"`
	InstallmentCount int       `json:"installment_count" validate:"required"`
	Interest         float64   `json:"interest" validate:"required"`
	AssetName        string    `json:"asset_name" validate:"required"`
	AssetType        string    `json:"asset_type" validate:"required"`
}

type TransactionResponse struct {
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
	CreatedAt        string    `json:"created_at"`
	UpdatedAt        string    `json:"updated_at"`
}
