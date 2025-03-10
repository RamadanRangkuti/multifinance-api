package transactions

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	model "github.com/RamadanRangkuti/multifinance-api/models"
	"github.com/google/uuid"
)

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *store {
	return &store{db: db}
}

type TransactionRepository interface {
	CreateTransaction(req model.TransactionRequest) (*model.TransactionResponse, error)
	GetAllTransactions(userID uuid.UUID) ([]model.TransactionResponse, error)
}

func (s *store) CreateTransaction(req model.TransactionRequest) (*model.TransactionResponse, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Ambil dan kunci limit pengguna
	var limitAmount, usedAmount float64
	queryLimit := `SELECT limit_amount, used_amount FROM user_limits WHERE user_id = $1 AND tenor = $2 FOR UPDATE`
	err = tx.QueryRow(queryLimit, req.UserID, req.Tenor).Scan(&limitAmount, &usedAmount)
	if err != nil {
		return nil, errors.New("failed to get user limit: " + err.Error())
	}

	// Cek apakah limit cukup
	if limitAmount-usedAmount < req.OTR {
		return nil, errors.New("insufficient limit")
	}

	// Generate contract number dengan MAX(contract_number)
	today := time.Now().Format("20060102") // Format YYYYMMDD
	var lastContractNumber string
	queryMaxContract := `SELECT contract_number FROM transactions WHERE DATE(created_at) = CURRENT_DATE ORDER BY contract_number DESC LIMIT 1`
	err = tx.QueryRow(queryMaxContract).Scan(&lastContractNumber)

	// Tentukan nomor urut berikutnya
	nextNumber := 1
	if err == nil && len(lastContractNumber) > 13 {
		var lastNumber int
		fmt.Sscanf(lastContractNumber[13:], "%d", &lastNumber)
		nextNumber = lastNumber + 1
	}

	contractNumber := fmt.Sprintf("TXN-%s-%04d", today, nextNumber)

	// Update used_amount di user_limits
	queryUpdateLimit := `UPDATE user_limits SET used_amount = used_amount + $1 WHERE user_id = $2 AND tenor = $3`
	_, err = tx.Exec(queryUpdateLimit, req.OTR, req.UserID, req.Tenor)
	if err != nil {
		return nil, errors.New("failed to update user limit: " + err.Error())
	}

	// Insert transaksi baru dengan contract_number yang sudah di-generate
	var transaction model.TransactionResponse
	queryTransaction := `
		INSERT INTO transactions (user_id, tenor, contract_number, otr, admin_fee, installment_count, interest, asset_name, asset_type, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW())
		RETURNING id, user_id, tenor, contract_number, otr, admin_fee, installment_count, interest, asset_name, asset_type, created_at, updated_at
	`
	err = tx.QueryRow(queryTransaction, req.UserID, req.Tenor, contractNumber, req.OTR, req.AdminFee, req.InstallmentCount, req.Interest, req.AssetName, req.AssetType).
		Scan(&transaction.ID, &transaction.UserID, &transaction.Tenor, &transaction.ContractNumber, &transaction.OTR, &transaction.AdminFee, &transaction.InstallmentCount, &transaction.Interest, &transaction.AssetName, &transaction.AssetType, &transaction.CreatedAt, &transaction.UpdatedAt)

	if err != nil {
		return nil, errors.New("failed to insert transaction: " + err.Error())
	}

	// Commit transaksi jika semua berhasil
	if err := tx.Commit(); err != nil {
		return nil, errors.New("failed to commit transaction: " + err.Error())
	}

	return &transaction, nil
}

func (s *store) GetAllTransactions(userID uuid.UUID) ([]model.TransactionResponse, error) {
	query := `SELECT id, user_id, tenor, contract_number, otr, admin_fee, installment_count, interest, asset_name, asset_type, created_at, updated_at 
	          FROM transactions WHERE user_id = $1`
	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.TransactionResponse
	for rows.Next() {
		var transaction model.TransactionResponse
		err := rows.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.Tenor,
			&transaction.ContractNumber,
			&transaction.OTR,
			&transaction.AdminFee,
			&transaction.InstallmentCount,
			&transaction.Interest,
			&transaction.AssetName,
			&transaction.AssetType,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
